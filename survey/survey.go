package survey

import (
	"fmt"
	"gopkg.in/AlecAivazis/survey.v1"
)

type model struct {
	Server string `survey:"server"`
	Port    int
}

var qs = []*survey.Question{
	{
		Name: "server",
		Prompt: &survey.Select{
			Message: "请选择一种服务状态: ",
			Options: []string{"tcp", "tcpserver", "udp", "udpserver"},
			Default: "tcp",
		},
	},
	{
		Name:   "port",
		Prompt: &survey.Input{Message: "请输入端口号"},
	},
}

func Start() model{
	answers := &model{}
	err := survey.Ask(qs, answers)
	if err != nil {
		fmt.Println(err.Error())
		return model{}
	}
	return *answers
}
