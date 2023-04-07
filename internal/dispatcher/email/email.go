package email

import (
	"context"
	"fmt"
	"github.com/summary/internal/core/domain"
	"github.com/summary/internal/core/port/outbound"
	"gopkg.in/gomail.v2"
	"log"
	"os"
	"time"
)

type service struct {
	m gomail.Dialer
}

func NewEmailService(m gomail.Dialer) outbound.DispatchMessageService {
	return &service{
		m: m,
	}
}

func (s *service) DispatchMessage(ctx context.Context, user domain.BalanceUser) error {
	log.Println("Sending email to user", user.EmailUser)
	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("ORIGIN_EMAIL"))
	m.SetHeader("To", user.EmailUser)
	m.SetHeader("Subject", "Heres is your summary!")
	m.SetBody("text/html", buildBody(user))

	if err := s.m.DialAndSend(m); err != nil {
		log.Println("Error sending email, error %s", err.Error())
		return err
	}
	return nil
}

func buildBody(user domain.BalanceUser) string {
	body := fmt.Sprintf("Hello !")
	body += fmt.Sprintf("<p>Here is your Balance summary <br>")
	body += fmt.Sprintf("Total Balance: <b>%.2f</b></p>", user.TotalBalance)
	body += fmt.Sprintf("Average credit amount: <b>%.2f</b><br>", user.TotalCreditBalance)
	body += fmt.Sprintf("Average debit amount: <b>%.2f</b><br>", user.TotalDebitBalance)
	body += fmt.Sprintf("<br><b>Month description</b><br>")
	return fillMonthAverage(user, body)
}

func fillMonthAverage(user domain.BalanceUser, body string) string {
	for i, v := range user.MonthTransactions {
		if v != 0 {
			body += fmt.Sprintf("<p> <b>%s</b>: %d transactions<br> <b>Average credit:</b> %.2f <br> <b>Average debit</b> %.2f </p>", time.Month(i+1).String(), v, user.MonthCreditBalance[i], user.MonthDebitBalance[i])
		}
	}
	return body
}
