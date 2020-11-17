package main

import (
	"fmt"
	"github.com/fatih/color"
	"handler"
	"net"
	"os"
	"requests"
	// _ "mybubletea"
)

const (
	tcpPort = ":8092"
	udpPort = 8091
	host    = "127.0.0.1"
)

func main() {
	var input string
	c := color.New(color.FgHiBlack, color.BgCyan)
	c.Printf("请选择是选择tcp, tcpserver, udp, udpserver之一: ")
	fmt.Scan(&input)
	// 判断选择运行的模式
	// tcp tcpserver udp udpserver
	switch input {
	case "tcp":
		// 创建conn
		conn, err := net.Dial("tcp", host+tcpPort)
		defer conn.Close()
		if err != nil {
			handler.Error("创建conn失败")
			return
		}
		// 输入要发送给服务器的内容
		fmt.Scan(&input)
		// 写入信息发送到服务端
		_, err = conn.Write([]byte(input))
		if err != nil {
			handler.Error("写入conn失败")
			return
		}
		// 准备接受服务端的信息
		buf := make([]byte, 1024)
		_, err = conn.Read(buf)
		if err != nil {
			handler.Error("接受不到服务端返回的信息")
		}
		handler.Recv(string(buf))
		break
	case "tcpserver":
		listener, err := net.Listen("tcp", tcpPort)
		if err != nil {
			handler.Error("创建listener失败")
			return
		}
		defer listener.Close()
		for {
			// 接受conn连接
			conn, err := listener.Accept()
			if err != nil {
				handler.Error("conn连接失败")
				return
			}
			// 创建处理tcp链接的groutine
			go tcpServerHandle(conn)
		}
	case "udp":
		conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
			IP:   net.IPv4(0, 0, 0, 0),
			Port: udpPort,
		})
		if err != nil {
			handler.Error("创建conn失败")
		}
		fmt.Scan(&input)
		conn.Write([]byte(input))
		break
	case "udpserver":
		listener, err := net.ListenUDP("udp", &net.UDPAddr{
			IP:   net.IPv4(0, 0, 0, 0),
			Port: udpPort,
		})
		if err != nil {
			handler.Error("udpserver创建失败")
		}
		for {
			recData := make([]byte, 1024)
			_, addr, err := listener.ReadFromUDP(recData)
			if err != nil {
				handler.Error("udpserver接收失败")
			}
			handler.Recv("来自ip为：" + addr.IP.String() + "消息为：" + string(recData))
		}
	default:
		handler.Error("输入错误")
		os.Exit(1)
		break
	}
}
func tcpServerHandle(conn net.Conn) {
	defer conn.Close()
	handler.Inform("客户端与服务器连接建立成功...")
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		handler.Error("读取错误")
		return
	}
	handler.Recv(string(buf[:n]))
	resp := requests.Get(string(buf[:n]))
	fmt.Println(resp)
	conn.Write([]byte(resp))
}
