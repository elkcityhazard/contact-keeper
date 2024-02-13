package mailer

import (
	"bytes"
	"embed"
	"html/template"
	"time"

	"github.com/elkcityhazard/contact-keeper/internal/config"
	"github.com/go-mail/mail/v2"
)

//go:embed "templates"
var templateFS embed.FS

var app *config.AppConfig

func NewMailerConfig(a *config.AppConfig) {
	app = a

}

type Mailer struct {
	dialer *mail.Dialer
	sender string
}

type Message struct {
	From        string
	To          string
	Subject     string
	PlainBody   string
	HTMLBody    string
	Attachments []string
}

func New(host string, port int, username, password, sender string) *Mailer {

	dialer := mail.NewDialer(host, port, username, password)
	dialer.Timeout = 5 * time.Second

	return &Mailer{
		dialer: dialer,
		sender: sender,
	}
}

func (m *Mailer) SendEmail(msg Message, errorChan <-chan error) {

	defer app.WG.Done()
	tmpl, err := template.New("email").ParseFS(templateFS, "templates/"+templateFile)

	if err != nil {
		app.ErrorChan <- err
	}

	emailSubject := new(bytes.Buffer)

	err = tmpl.ExecuteTemplate(emailSubject, "subject", msg.Subject)
	if err != nil {
		app.ErrorChan <- err
	}

	plainBody := new(bytes.Buffer)

	err = tmpl.ExecuteTemplate(plainBody, "plainBody", msg.PlainBody)
	if err != nil {
		app.ErrorChan <- err
	}

	htmlBody := new(bytes.Buffer)

	err = tmpl.ExecuteTemplate(htmlBody, "htmlBody", msg.HTMLBody)
	if err != nil {
		app.ErrorChan <- err
	}

	mailMsg := mail.NewMessage()
	mailMsg.SetHeader("To", msg.To)
	mailMsg.SetHeader("From", msg.From)
	mailMsg.SetHeader("Subject", emailSubject.String())
	mailMsg.SetBody("text/plain", plainBody.String())
	mailMsg.AddAlternative("text/html", htmlBody.String())

	err = m.dialer.DialAndSend(mailMsg)

	if err != nil {
		app.ErrorChan <- err
	}
}

func (m *Mailer) ListenForMail() {
	for {
		select {
		case msg := <-app.MailerChan:
			go m.SendEmail(msg, app.ErrorChan)
		case <-app.ErrorChan:
			return
		}
	}
}
