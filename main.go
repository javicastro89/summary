package main

import (
	"github.com/summary/internal/factory"
	"log"
)

func main() {
	f := factory.New()
	if err := f.Run(); err != nil {
		log.Fatal("Error running the app. Error %s", err.Error())
	}
	//_ = os.Setenv("_LAMBDA_SERVER_PORT", "587")
	//_ = os.Setenv("AWS_LAMBDA_RUNTIME_API", "provided.al2")
	//lambda.Start(handler)

	//}

	//func handler() {
	//cfg, err := config.LoadDefaultConfig(context.Background())
	//if err != nil {
	//	panic(err)
	//}
	//ses := sesv2.NewFromConfig(cfg)

	//body := "Hello <b>Bob</b>!"
	//subject := "The subject"
	//arn := "arn:aws:ses:us-east-1:658012167973:identity/no.reply.summary.test@gmail.com"

	//email, err := ses.SendEmail(context.Background(), &sesv2.SendEmailInput{
	//	Content: &types.EmailContent{
	//		Simple: &types.Message{
	//			Body: &types.Body{
	//				Html: &types.Content{
	//					Data: &body,
	//				},
	//				Text: &types.Content{
	//					Data: &body,
	//				},
	//			},
	//			Subject: &types.Content{
	//				Data: &subject,
	//			},
	//		},
	//	},
	//	//ConfigurationSetName:        nil,
	//	Destination: &types.Destination{
	//		ToAddresses: []string{"javiercastroloureiro@gmail.com"},
	//	},
	//	FromEmailAddressIdentityArn: &arn,
	//})
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println("EmailId", email.MessageId)

	// Send the email to Bob

	//m := gomail.NewMessage()
	//m.SetHeader("From", "no.reply.summary.test@gmail.com")
	//m.SetHeader("To", "javiercastroloureiro@gmail.com")
	//m.SetHeader("Subject", "Hello!")
	//m.SetBody("text/html", "Hello <b>Bob</b>!")
	//d := gomail.NewDialer("email-smtp.us-east-1.amazonaws.com", 587, "AKIAZSNEU34S7MJPUETR", "BGAgBv/by3FblCpB2RaXlcuocjQXF7s8QqaRhvVVGs3A")
	//if err := d.DialAndSend(m); err != nil {
	//	panic(err)
	//}
}
