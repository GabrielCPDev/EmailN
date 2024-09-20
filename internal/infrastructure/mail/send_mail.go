package mail

import (
	"fmt"
	"os"

	"gopkg.in/gomail.v2"
)

func SendMail() error{
	fmt.Println("enviando email ...")

	d := gomail.NewDialer(os.Getenv("EMAIL_SMTP"), 587, os.Getenv("EMAIL_USER"), os.Getenv("EMAIL_PASSWORD"))

	m := gomail.NewMessage()
	m.SetHeader("from", os.Getenv("EMAIL_USER"))
	m.SetHeader("To", os.Getenv("user@user.com"))
	m.SetHeader("subject", os.Getenv("Hello!"))
	m.SetBody("text/html", "Hello <b> Satoshi </>!")

	return d.DialAndSend(m)
}