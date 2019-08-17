package main

import (
	"net/smtp"
	"strings"
	"fmt"
)

const (
	HOST = "smtp.163.com"
	SERVER_ADDR = "smtp.163.com:25"
	USER = "dxb350352@163.com" //发送邮件的邮箱
	PASSWORD = "1qaz@WSX"         //发送邮件邮箱的密码
)

type Email struct {
	to      string "to"
	subject string "subject"
	msg     string "msg"
}

func NewEmail(to, subject, msg string) *Email {
	return &Email{to: to, subject: subject, msg: msg}
}

func SendEmail(email *Email) error {
	auth := smtp.PlainAuth("", USER, PASSWORD, HOST)
	sendTo := strings.Split(email.to, ";")

	for _, v := range sendTo {
		str := strings.Replace("From: " + USER + "~To: " + v + "~Subject: " + email.subject + "~~", "~", "\r\n", -1) + email.msg
		fmt.Println(str)
		err := smtp.SendMail(
			SERVER_ADDR,
			auth,
			USER,
			[]string{v},
			[]byte(str),
		)
		if (err != nil) {
			return err
		}
	}

	return nil
}

func main() {
	//fmt.Println(SendEmail(NewEmail("542724976@qq.com", "subj11ect", "messages")))
	test263();
}

func test263() {
	auth := smtp.PlainAuth("", "ceshi1@wondersoft.cn", "pengwenming123", "smtp.263.net")
	err := smtp.SendMail(
		"smtp.263.net:25",
		auth,
		"ceshi1@wondersoft.cn",
		[]string{"542724976@qq.com","dxb350352@163.com"},
		[]byte("messages"),
	)
	fmt.Println(err)
}
