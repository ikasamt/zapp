package zapp

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sendgrid/rest"
	sendgrid "github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/urlfetch"
)

func SendEmailGAE(c *gin.Context, sendgridAPIKey string, from string, tos []string, subject string, body string) error {
	ctx := appengine.NewContext(c.Request)

	// 送信
	from_ := mail.NewEmail("", from)

	// 送信
	sendgrid.DefaultClient = &rest.Client{HTTPClient: urlfetch.Client(ctx)}
	client := sendgrid.NewSendClient(sendgridAPIKey)
	plainText := mail.NewContent("text/plain", body)

	m := new(mail.SGMailV3)
	m.SetFrom(from_)
	m.Subject = subject
	p := mail.NewPersonalization()
	for _, toStr := range tos {
		to := mail.NewEmail("", toStr)
		p.AddTos(to)
	}
	m.AddPersonalizations(p)
	m.AddContent(plainText)

	res, err := client.Send(m)
	if err != nil {
		log.Errorf(ctx, "send mail failed / err: %v", err)
		return err
	}

	if res.StatusCode >= 400 {
		log.Errorf(ctx, "send mail failed / code %d: %v", res.StatusCode, res.Body)
		return fmt.Errorf(res.Body)
	}

	return nil
}
