package handler

import (
	"github.com/fatih/color"
)

var c *color.Color

func init(){
}


func Error(message string){
  c = color.New(color.FgHiWhite, color.BgRed)
  c.Println("发生错误: " + message)
}


func Recv(message string){
  c = color.New(color.FgHiWhite, color.BgBlue)
  c.Println("受到信息: " + message)
}

func Send(message string){
  c = color.New(color.FgHiWhite, color.BgGreen)
  c.Println("发送信息" + message)
}


func Inform(message string){
  c = color.New(color.FgHiBlack, color.BgYellow)
  c.Println(message)
}
