package factory

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	db "github.com/aws/aws-sdk-go-v2/service/dynamodb"
	s3Client "github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/summary/internal/controller/s3"
	"github.com/summary/internal/core/port/inbound"
	"github.com/summary/internal/core/port/outbound"
	"github.com/summary/internal/core/service"
	"github.com/summary/internal/dispatcher/email"
	"github.com/summary/internal/repository/dynamodb"
	"gopkg.in/gomail.v2"
	"log"
	"os"
	"strconv"
)

type Factory struct {
	srv        inbound.ProcessSummaryService
	dispatch   outbound.DispatchMessageService
	repository outbound.RepositoryService
	emailSrv   *gomail.Dialer
	s3Client   *s3Client.Client
}

func New() *Factory {
	return &Factory{}
}

func (f *Factory) Run() error {
	log.Printf("Starting factory")
	if err := f.setEmailService(); err != nil {
		log.Fatalf("Error seting email %s", err.Error())
		return err
	}
	if err := f.setS3(); err != nil {
		log.Fatalf("Error seting s3 %s", err.Error())
		return err
	}
	if err := f.setRepository(); err != nil {
		log.Fatalf("Error seting repository %s", err.Error())
		return err
	}
	f.setDispatch()
	f.setService()
	f.run()
	log.Printf("End factory")
	return nil
}

// d := gomail.NewDialer("email-smtp.us-east-1.amazonaws.com", 587, "AKIAZSNEU34S7MJPUETR", "BGAgBv/by3FblCpB2RaXlcuocjQXF7s8QqaRhvVVGs3A")
func (f *Factory) setEmailService() error {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatalf("Error Atoi port. Error: %s", err.Error())
		return err
	}
	f.emailSrv = gomail.NewDialer(os.Getenv("HOST"), port, os.Getenv("USERNAME"), os.Getenv("PASSWORD"))
	return nil
}

func (f *Factory) setS3() error {
	sdkConfig, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		log.Printf("failed to load default config: %s", err)
		return err
	}
	f.s3Client = s3Client.NewFromConfig(sdkConfig)
	return nil
}

func (f *Factory) setDispatch() {
	f.dispatch = email.NewEmailService(*f.emailSrv)
}

func (f *Factory) setRepository() error {
	conf, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
		return err
	}
	dbClient := *db.NewFromConfig(conf)

	f.repository = dynamodb.NewRepositoryService(dbClient, os.Getenv("TABLE_NAME"))
	return nil
}

func (f *Factory) setService() {
	f.srv = service.NewService(f.dispatch, f.repository)
}

func (f *Factory) run() {
	ctrl := s3.NewController(f.srv, f.s3Client)
	log.Println("Starting handler with ctrl: ", ctrl)
	lambda.Start(ctrl.Handler)
}
