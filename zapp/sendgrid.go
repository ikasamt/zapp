package zapp

import (
	"fmt"
	"log"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendEmail(sendgridAPIKey string, from string, to string, subject string, body string) error {

	// 送信
	from_ := mail.NewEmail("", from)
	to_ := mail.NewEmail("", to)

	// 送信
	client := sendgrid.NewSendClient(sendgridAPIKey)
	plainText := mail.NewContent("text/plain", body)
	message := mail.NewV3MailInit(from_, subject, to_, plainText)
	// response, err := client.Send(message)
	_, err := client.Send(message)
	if err != nil {
		log.Println(err)
		return err
	}
	// log.Println(`--------`)
	// log.Println("status", response.StatusCode)
	// log.Println("body", response.Body)
	// log.Println("headers", response.Headers)
	// log.Println(`--------`)
	return nil
}

func SendEmailMultiTo(sendgridAPIKey string, from string, toStrings []string, subject string, body string) error {
	// 送信
	client := sendgrid.NewSendClient(sendgridAPIKey)
	plainText := mail.NewContent("text/plain", body)

	// message
	m := new(mail.SGMailV3)
	from_ := mail.NewEmail("", from)
	m.SetFrom(from_)
	m.Subject = subject
	m.AddContent(plainText)

	p := mail.NewPersonalization()
	tos := []*mail.Email{}
	for _, toString := range toStrings {
		tos = append(tos, mail.NewEmail("", toString))
	}
	p.AddTos(tos...)
	m.AddPersonalizations(p)

	res, err := client.Send(m)
	if err != nil {
		errorMsg := fmt.Sprintf("send mail failed / err: %v", err)
		log.Println(errorMsg)
		return err
	}

	if res.StatusCode >= 400 {
		errorMsg := fmt.Sprintf("send mail failed / code %d: %v", res.StatusCode, res.Body)
		log.Println(errorMsg)
		return fmt.Errorf(res.Body)
	}

	return nil
}
