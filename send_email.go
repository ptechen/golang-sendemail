package sendemail

import (
	"fmt"
	"net/smtp"
	"strings"
)

type unencryptedAuth struct {
	smtp.Auth
}

func (a *unencryptedAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	s := *server
	s.TLS = true
	return a.Auth.Start(&s)
}

type SendEmailParams struct {
	UserEmail string
	Password string
	Host string
	Port int
	ContentType string  //"Content-Type: text/plain; charset=UTF-8"
} 

//通过未加密的连接发送电子邮件
func (self *SendEmailParams) UnencryptedSendEmail(sendMsg, subject, nickName string, send2Email []string) (err error) {
	auth := &unencryptedAuth{
		smtp.PlainAuth("", self.UserEmail, self.Password, self.Host),
	}
	msgStr := fmt.Sprintf("To: %s \r\nFrom: %s <%s>\r\nSubject: %s\r\n%s\r\n\r\n%s",
		strings.Join(send2Email, ","), nickName, self.UserEmail, subject, self.ContentType, sendMsg)
	msgByte := []byte(msgStr)
	address := fmt.Sprintf("%s:%d", self.Host, self.Port)
	err = smtp.SendMail(address, auth, self.UserEmail, send2Email, msgByte)
	return
}

//通过加密的连接发送电子邮件
func (self *SendEmailParams) SendEmail(sendMsg, subject, nickName string, send2Email []string) (err error) {
	auth := smtp.PlainAuth("", self.UserEmail, self.Password, self.Host)
	msgStr := fmt.Sprintf("To: %s \r\nFrom: %s <%s>\r\nSubject: %s\r\n%s\r\n\r\n%s",
		strings.Join(send2Email, ","), nickName, self.UserEmail, subject, self.ContentType, sendMsg)
	msgByte := []byte(msgStr)
	address := fmt.Sprintf("%s:%d", self.Host, self.Port)
	err = smtp.SendMail(address, auth, self.UserEmail, send2Email, msgByte)
	return
}

