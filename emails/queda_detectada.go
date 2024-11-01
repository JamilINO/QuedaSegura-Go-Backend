package mypkgteste

import (
	"log"
	"os"
	"github.com/wneessen/go-mail"
)

func Send()  {
    user := os.Getenv("EMAIL_USER")

    message := mail.NewMsg()
	if err := message.From(user); err != nil {
		log.Fatalf("failed to set From address: %s", err)
	}
	if err := message.To(user); err != nil {
		log.Fatalf("failed to set To address: %s", err)
	}

    message.Subject("This is my first mail with go-mail!")
	message.SetBodyString(mail.TypeTextPlain, "Do you like this mail? I certainly do!")
   
	client, _ := mail.NewClient("smtp.gmail.com",
		mail.WithSMTPAuth(mail.SMTPAuthPlain), mail.WithTLSPortPolicy(mail.TLSMandatory),
		mail.WithUsername(user), mail.WithPassword(os.Getenv("EMAIL_PASS")),
	)

    if err := client.DialAndSend(message); err != nil {
		log.Fatalf("failed to send mail: %s", err)
	}	
}