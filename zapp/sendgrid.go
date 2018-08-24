package zapp

import (
	"log"

	sendgrid "github.com/sendgrid/sendgrid-go"
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
