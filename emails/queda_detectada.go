package emails

import (
	"context"
	//"encoding/base64"
	"fmt"
	"html/template"
	"log"
	"os"

	"github.com/wneessen/go-mail"
	"quedasegura.com/m/v2/db"
	"quedasegura.com/m/v2/proto/convert"
)

type HtmlProps struct{
	info convert.QuedaPayload;
	logo string;
}


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

	var email string

	rows, db_err := db.Postgres.Query(context.Background(), `
	SELECT Contacts.email FROM Devices 
	INNER JOIN Users ON Users.id = Devices.foreign_id 
	INNER JOIN Contacts ON Users.id = Contacts.foreign_id
	WHERE mac_adress = $1;
	`, info.MacAddr)

	for rows.Next() {
		rows.Scan(&email)
		if err := message.AddTo(email); err != nil {
			log.Fatalf("failed to set To address: %s", err)
		}
	}

	if db_err != nil {
		fmt.Errorf("Erro no DB worker")
		return
	}

	//fmt.Print("\n\n\nEntered Here\n\n\n")

	target_time := info.Time.AsTime()
    message.Subject(fmt.Sprintf("Alerta de Queda às %d:%d:%d no dia %d/%d/%d", target_time.Hour(),target_time.Minute(), target_time.Second(), target_time.Day(), target_time.Month(), target_time.Year()))

	template_file, err := os.ReadFile("./assets/queda_template.html")
	
	if err != nil {
		fmt.Printf(err.Error())
	}

	img, err := os.ReadFile("./assets/logo.png")

	if err != nil {
		fmt.Printf(err.Error())
	}

	html, err := template.New("htmltpl").Parse(string(template_file),)

	if err != nil {
		fmt.Printf(err.Error())
	}

	//base_img

	fmt.Printf(string(img))

	tmpl_err := message.SetBodyHTMLTemplate(html, nil)
   
	//message.AttachFile("./assets/logo.png")

	if tmpl_err != nil {
		fmt.Printf(err.Error())
	}


	client, _ := mail.NewClient(mail_server,
		mail.WithSMTPAuth(mail.SMTPAuthPlain), mail.WithTLSPortPolicy(mail.TLSMandatory),
		mail.WithUsername(user), mail.WithPassword(os.Getenv("EMAIL_PASS")),
	)

    if err := client.DialAndSend(message); err != nil {
		log.Fatalf("failed to send mail: %s", err)
	}	
	
}