package s3

import (
	"github.com/summary/internal/core/domain"
	"log"
	"strconv"
	"strings"
)

type User struct {
	Id          string `csv:"Id"`
	Date        string `csv:"Date"`
	Transaction string `csv:"Transaction"`
}

func (u User) toUserSummary() (*domain.User, error) {
	strNum := strings.ReplaceAll(u.Transaction, ",", ".")
	parse, err := strconv.ParseFloat(strNum, 64)
	if err != nil {
		log.Fatalf("Error parsing float. Error: %s", err.Error())
		return nil, err
	}
	return &domain.User{
		Id:          u.Id,
		Date:        u.Date,
		Transaction: parse,
	}, nil
}
