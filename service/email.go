package service

import (
	"fmt"
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
	body := fmt.Sprintln(time.Now().Date())
	return controller.SendMail(mailTo, subject, body)
}
