package controller

import (
	"strconv"

	"github.com/AllinChen/systool01/config"
	"gopkg.in/gomail.v2"
)

/*
	//定义收件人
	mailTo := []string{
		config.GlobalConf.Email.Toemail,
	}
	//邮件主题为"Hello"
	subject := "每日运行统计"
	// 邮件正文
	body := fmt.Sprintln(time.Now().Date())
	controller.SendMail(mailTo, subject, body)
*/
func SendMail(mailTo []string, subject string, body string) error {
	mailConn := map[string]string{
		"user": config.GlobalConf.Email.User,
		"pass": config.GlobalConf.Email.Pass,
		"host": config.GlobalConf.Email.Host,
		"port": config.GlobalConf.Email.Port,
	}

	port, _ := strconv.Atoi(mailConn["port"]) //转换端口类型为int

	m := gomail.NewMessage()
	m.SetHeader("From", "SysTool01"+"<"+mailConn["user"]+">") //这种方式可以添加别名，即“XD Game”， 也可以直接用<code>m.SetHeader("From",mailConn["user"])</code> 读者可以自行实验下效果
	m.SetHeader("To", mailTo...)                              //发送给多个用户
	m.SetHeader("Subject", subject)                           //设置邮件主题
	m.SetBody("text/html", body)                              //设置邮件正文

	d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])

	err := d.DialAndSend(m)
	return err

}
