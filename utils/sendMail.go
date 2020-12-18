package utils

import (
	"encoding/json"
	"fmt"
	"gopkg.in/gomail.v2"
	"mime"
	"strconv"
	"sync"
)

type Msg struct {
	Code    int    `json:"code"`
	Mail    string `json:"mail"`
	Message string `json:"message"`
}

var res Msg

// gomail -u 2387986432@qq.com -auth 1234 -p 587 -h smtp.qq.com -sub title -m message -from 123@163.com
func SendMail(account map[string]string, mailTo string, subject, body string, wg *sync.WaitGroup) {
	defer wg.Done()
	port, _ := strconv.Atoi(account["port"])
	msg := gomail.NewMessage()
	msg.SetHeader("From", mime.QEncoding.Encode("UTF-8", "<"+account["username"]+">"))
	msg.SetHeader("To", mailTo)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", body)
	d := gomail.NewDialer(account["host"], port, account["username"], account["authCode"])
	err := d.DialAndSend(msg)
	if err != nil {
		res.Code = 0
		res.Mail = mailTo
		res.Message = err.Error()
	} else {
		res.Code = 1
		res.Mail = mailTo
		res.Message = "Send Email Successfully"
	}
	result, _ := json.Marshal(res)
	fmt.Println(string(result))
}
