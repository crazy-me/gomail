package main

import (
	"flag"
	"github.com/crazy-me/gomail/utils"
	"strings"
	"sync"
)

func main() {
	var username string
	var authCode string
	var port string
	var host string
	var subject string
	var body string
	var fromTo string

	flag.StringVar(&username, "u", "", "发送邮箱账号")
	flag.StringVar(&authCode, "auth", "", "邮箱授权码")
	flag.StringVar(&port, "p", "", "端口")
	flag.StringVar(&host, "h", "", "SMTP地址")
	flag.StringVar(&subject, "sub", "", "邮件主题")
	flag.StringVar(&body, "m", "", "邮件正文")
	flag.StringVar(&fromTo, "from", "", "接收者邮箱地址,多个以逗号分隔")

	// 解析命令行参数
	flag.Parse()

	if username == "" || authCode == "" || port == "" || host == "" || subject == "" || body == "" || fromTo == "" {
		flag.PrintDefaults()
		return
	}

	mailSlide := strings.Split(fromTo, ",")

	account := map[string]string{
		"username": username,
		"authCode": authCode,
		"host":     host,
		"port":     port,
	}
	var wg sync.WaitGroup
	for _, mail := range mailSlide {
		wg.Add(1)
		go utils.SendMail(account, mail, subject, body, &wg)
	}
	wg.Wait()
}
