package main

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

const (
	userName    = "***@***.com"
	passWord    = "******"
	host        = "***.***.com"
	contentType = "Content-Type: text/plain; charset=UTF-8"
	port        = 25
)

func SendEmail(sendMsg, subject, nickName string, send2Email []string) (err error) {
	auth := &unencryptedAuth{
		smtp.PlainAuth("", userName, passWord, host),
	}
	msgStr := fmt.Sprintf("To: %s \r\nFrom: %s <%s>\r\nSubject: %s\r\n%s\r\n\r\n%s",
		strings.Join(send2Email, ","), nickName, userName, subject, contentType, sendMsg)
	msgByte := []byte(msgStr)
	address := fmt.Sprintf("%s:%d", host, port)
	err = smtp.SendMail(address, auth, userName, send2Email, msgByte)
	return
}

func main() {
	// Set up authentication information.
	sendMsg := "SendEmail"
	nickName := "testok"
	subject := "testok mail"
	send2Email := []string{"***@***.com", "***@***.com"}
	err := SendEmail(sendMsg, subject, nickName, send2Email)
	fmt.Println(err)
}

