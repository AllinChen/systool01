package service

import (
	"fmt"
	"strings"
	"time"

	"github.com/AllinChen/systool01/config"
	"github.com/AllinChen/systool01/controller"
)

func SendMail() error {
	//定义收件人
	mailTo := []string{
		// config.GlobalConf.Email.Toemail,
		config.GlobalConf.Email.User,
	}
	//邮件主题为"Hello"
	subject := "每日运行统计"
	// 邮件正文
	body := fmt.Sprintln(time.Now().Date()) + "<br>" + MakeSysInfoMail()
	fmt.Println(body)
	return controller.SendMail(mailTo, subject, body)
}

func MakeSysInfoMail() (ms string) {
	_, ms = controller.GetSysRunInfo()
	ms = strings.Replace(ms, "\n", "<br>", -1)

	return
}
