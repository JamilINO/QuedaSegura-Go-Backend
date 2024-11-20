package emails

import (
	"context"

	"encoding/base64"
	"fmt"
	"html/template"
	"log"
	"os"

	"github.com/wneessen/go-mail"
	"quedasegura.com/m/v2/db"
	"quedasegura.com/m/v2/proto/convert"
)

type HtmlProps struct{
	EmailList string;
	Time string;
	Date string;

	MacAddr string;
	Name string

	Logo string;
	Url string;
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

	if db_err != nil{
		fmt.Printf(db_err.Error())
		return
	}

	var email_list string = ""
	for rows.Next() {
		rows.Scan(&email)
		email_list += email + " "
		if err := message.AddTo(email); err != nil {
			log.Fatalf("failed to set To address: %s", err)
		}
	}

	if email_list == ""{
		fmt.Printf("Email nil")
		return
	}

	if db_err != nil {
		fmt.Errorf("Erro no DB worker")
		return
	}
	
	var real_name string
	err := db.Postgres.QueryRow(context.Background(), `
	SELECT Users.real_name FROM Users
	INNER JOIN Devices ON Users.id = Devices.foreign_id 
	WHERE mac_adress=$1 LIMIT 1;
	`, info.MacAddr).Scan(&real_name)

	fmt.Printf("%s", real_name)

	if err != nil {
		fmt.Sprintf(err.Error())
	}

	//fmt.Print("\n\n\nEntered Here\n\n\n")

	target_time := info.Time.AsTime()
	fmt_time := fmt.Sprintf("%d:%d:%d", target_time.Hour(),target_time.Minute(), target_time.Second())
	fmt_day := fmt.Sprintf("%d/%d/%d", target_time.Day(), target_time.Month(), target_time.Year())
    message.Subject(fmt.Sprintf("Alerta de Queda às %s no dia %s", fmt_time , fmt_day))

	template_file, err := os.ReadFile("./assets/queda_template.html")
	
	if err != nil {
		fmt.Printf(err.Error())
	}

	img, err := os.ReadFile("./assets/logo.png")

	if err != nil {
		fmt.Printf(err.Error())
	}

	html, err := template.New("htmltpl").Parse(string(template_file))

	if err != nil {
		fmt.Printf(err.Error())
	}

	//fmt.Printf(string(img))

	//r := fmt.Sprintf(info.MacAddr)


	
	err = message.SetBodyHTMLTemplate(html, HtmlProps{
		EmailList: email_list,
		Time: fmt_time,
		Date: fmt_day,

		MacAddr: info.MacAddr,
		Name: real_name,

		Logo: base64.StdEncoding.EncodeToString(img),
		Url: "http://192.168.15.77:7777",
	})

	if err != nil{
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