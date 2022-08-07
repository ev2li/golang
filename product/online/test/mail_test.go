package test

import (
	"crypto/tls"
	"net/smtp"
	"testing"

	"github.com/jordan-wright/email"
)

func TestSendMail(t *testing.T) {
	e := email.NewEmail()
	e.From = "Get <873677408@qq.com>"
	e.To = []string{"873677408@qq.com"}
	e.Subject = "验证码发送测试"
	e.HTML = []byte("<h1>Fancy HTML is supported, too!</h1>")
	// err := e.Send("smtp.163.com:465", smtp.PlainAuth("", "873677408@qq.com", "password123", "smtp.qq.com"))
	//返回EOF时，关闭SSL重试
	err := e.SendWithTLS("smtp.163.com:465", smtp.PlainAuth("", "873677408@qq.com", "password123", "smtp.qq.com"), &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         "smtp.163.com:465",
	})

	if err != nil {
		t.Fatal(err)
	}
}
