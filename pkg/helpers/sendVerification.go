package helpers

import (
	"bytes"
	"fmt"
	"go-restapi-boilerplate/models"
	"html/template"
	"log"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

// mail verification
func SendVerificationEmail(token string, user *models.MstUser) {
	var CONFIG_SMTP_HOST = os.Getenv("CONFIG_SMTP_HOST")
	var CONFIG_SMTP_PORT, _ = strconv.Atoi(os.Getenv("CONFIG_SMTP_PORT"))
	var CONFIG_SENDER_NAME = os.Getenv("CONFIG_SENDER_NAME")
	var CONFIG_AUTH_EMAIL = os.Getenv("CONFIG_AUTH_EMAIL")
	var CONFIG_AUTH_PASSWORD = os.Getenv("CONFIG_AUTH_PASSWORD")

	data := map[string]string{
		"Name":    user.FullName,
		"AppName": os.Getenv("APP_NAME"),
		"TOKEN":   token,
	}

	// mengambil file template
	t, err := template.ParseFiles("views/verificationEmail.html")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bodyMail := new(bytes.Buffer)

	// mengeksekusi template, dan memparse "data" ke template
	t.Execute(bodyMail, data)

	// create new message
	verificationEmail := gomail.NewMessage()
	verificationEmail.SetHeader("From", CONFIG_SENDER_NAME)
	verificationEmail.SetHeader("To", user.Email)
	verificationEmail.SetHeader("Subject", "Email Verification")
	verificationEmail.SetBody("text/html", bodyMail.String())

	verificationDialer := gomail.NewDialer(
		CONFIG_SMTP_HOST,
		CONFIG_SMTP_PORT,
		CONFIG_AUTH_EMAIL,
		CONFIG_AUTH_PASSWORD,
	)

	err = verificationDialer.DialAndSend(verificationEmail)
	if err != nil {
		fmt.Println("Gagal mengirim email")
		fmt.Println(err.Error())
		return
	}

	log.Println("Pesan terkirim!")
}
