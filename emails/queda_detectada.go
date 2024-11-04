package emails

import (
	"fmt"
	"log"
	"os"

	"github.com/wneessen/go-mail"
	"quedasegura.com/m/v2/proto/convert"
)

func Send(info *convert.QuedaPayload)  {
    user := os.Getenv("EMAIL_USER")
	mail_server := os.Getenv("EMAIL_SERVER")

	if mail_server == "" {
		mail_server = "smtp.gmail.com"
	}

    message := mail.NewMsg()
	if err := message.From(user); err != nil {
		log.Fatalf("failed to set From address: %s", err)
	}
	if err := message.To(user); err != nil {
		log.Fatalf("failed to set To address: %s", err)
	}

    message.Subject("This is my first mail with go-mail!")
	str := fmt.Sprintf("MacAddr: %s\nDate: %s\nIntensity: %.2f", info.MacAddr, info.Time.AsTime(), info.Intensity)
	message.SetBodyString(mail.TypeTextPlain, str)
   
	client, _ := mail.NewClient(mail_server,
		mail.WithSMTPAuth(mail.SMTPAuthPlain), mail.WithTLSPortPolicy(mail.TLSMandatory),
		mail.WithUsername(user), mail.WithPassword(os.Getenv("EMAIL_PASS")),
	)

    if err := client.DialAndSend(message); err != nil {
		log.Fatalf("failed to send mail: %s", err)
	}	
}