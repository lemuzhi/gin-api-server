package tools

import (
	"fmt"
	"gin-project-template/conf"
	"github.com/jordan-wright/email"
	"log"
	"net/smtp"
)

func SendEmailWarn(serverName, logid string) {
	title := fmt.Sprintf("%s服务发生异常，日志ID：%s", serverName, logid)

	e := email.NewEmail()
	e.From = conf.Email.From // 发送方
	e.To = conf.Email.To     //接收方
	e.Subject = title        //邮件标题
	e.HTML = []byte(title)   //邮件内容
	err := e.Send(conf.Email.Addr, smtp.PlainAuth("", conf.Email.From, conf.Email.Password, conf.Email.Host))
	if err != nil {
		log.Println(err)
	}
}
