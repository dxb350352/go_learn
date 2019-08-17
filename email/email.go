package main

import (
	"fmt"
	"github.com/jordan-wright/email"
	"net/smtp"
	"github.com/sas/utils"
)

func main() {
	err := utils.SendMail("smtp.163.com", "25", "wondersoft_email@163.com", "wondersoft123", "542724976@qq.com", "为什么", "为什么")
	//err := utils.SendMail("smtp.263.net", "25", "ceshi1@wondersoft.cn", "pengwenming123", "542724976@qq.com", "subj11ect", "message")
	//err := utils.SendMail("smtp.263.net", "25", "ceshi1@wondersoft.cn", "pengwenming123", "daixiaobo@wondersoft.cn", "subj11ect", "message")
	fmt.Println(err)
}
func SendMail(smtpserver, port, user, password, recipient, subject, message string) error {
	e := email.NewEmail()
	e.From = user
	e.To = []string{recipient}
	e.Subject = subject
	e.Text = []byte(message)
	err := e.Send(smtpserver + ":" + port, smtp.PlainAuth("", user, password, smtpserver))
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("go")
	return nil
}
