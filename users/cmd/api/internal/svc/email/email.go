package email

import (
	"crypto/tls"
	"gopkg.in/gomail.v2"
)

type Email struct {
	*SMTPInfo
}

// SMTPInfo 结构体用于传递发送邮箱所必需的信息
type SMTPInfo struct {
	Host     string
	Port     int
	IsSSL    bool
	UserName string
	Password string
	From     string
}

func NewEmail(info *SMTPInfo) *Email {
	return &Email{SMTPInfo: info}
}

func (e *Email) SendMail(to string, subject, body string) error {
	// NewMessage 方法创建一个消息实例,用于设置邮件的一些必要信息
	m := gomail.NewMessage()
	m.SetHeader("From", e.From)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	// NewDialer 方法创建一个新的 SMTP 拨号实例
	dialer := gomail.NewDialer(e.Host, e.Port, e.UserName, e.Password)
	//设置对应的拨号信息用于连接 SMTP 服务器
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: e.IsSSL}
	// DialAndSend 方法打开与 SMTP 服务器的连接并发送电子邮件
	return dialer.DialAndSend(m)
}
