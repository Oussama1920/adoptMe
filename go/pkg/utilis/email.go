package utils

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"

	db "github.com/Oussama1920/adoptMe/go/pkg/db"

	config "github.com/Oussama1920/adoptMe/go/pkg/config"
	"github.com/k3a/html2text"
	"gopkg.in/gomail.v2"
)

type EmailConfig struct {
	EMAIL_FROM string
	SMTP_HOST  string
	SMTP_USER  string
	SMTP_PASS  string
	SMTP_PORT  int
}
type EmailData struct {
	URL       string
	FirstName string
	Subject   string
}

// ? Email template parser

func ParseTemplateDir(dir string) (*template.Template, error) {
	var paths []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return template.ParseFiles(paths...)
}

func SendEmail(user db.User, data *EmailData) {
	var emailConfig EmailConfig
	if err := config.GetDataConfiguration("service.email", &emailConfig); err != nil {
		log.Fatal("could not load config", err)
	}

	// Sender data.
	from := emailConfig.EMAIL_FROM
	smtpPass := emailConfig.SMTP_PASS
	smtpUser := emailConfig.SMTP_USER
	to := user.Email
	smtpHost := emailConfig.SMTP_HOST
	smtpPort := emailConfig.SMTP_PORT

	var body bytes.Buffer

	template, err := ParseTemplateDir("templates")
	if err != nil {
		log.Fatal("Could not parse template", err)
	}

	template.ExecuteTemplate(&body, "verificationCode.html", &data)

	m := gomail.NewMessage()

	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", data.Subject)
	m.SetBody("text/html", body.String())
	m.AddAlternative("text/plain", html2text.HTML2Text(body.String()))
	fmt.Println("host : ", smtpHost)
	fmt.Println("port : ", smtpPort)
	fmt.Println("User : ", smtpUser)
	fmt.Println("Password : ", smtpPass)

	d := gomail.NewDialer(smtpHost, smtpPort, smtpUser, smtpPass)
	//	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send Email
	if err := d.DialAndSend(m); err != nil {
		log.Printf("Could not send email: %v ", err)
	}

}
