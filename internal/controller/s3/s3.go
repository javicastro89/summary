package s3

import (
	"context"
	"encoding/csv"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/jszwec/csvutil"
	"github.com/summary/internal/core/domain"
	"github.com/summary/internal/core/port/inbound"
	"io"
	"log"
)

type Controller struct {
	srv inbound.ProcessSummaryService
	s3  *s3.Client
}

func NewController(srv inbound.ProcessSummaryService, s3 *s3.Client) *Controller {
	return &Controller{
		srv: srv,
		s3:  s3,
	}
}

func (c *Controller) Handler(ctx context.Context, s3Event events.S3Event) error {
	log.Println("Processing records", s3Event.Records)
	for _, record := range s3Event.Records {
		objectOutput, errGettingObj := c.getS3Object(ctx, record)
		if errGettingObj != nil {
			log.Fatalf("error getting object %s/%s: %s", record.S3.Bucket.Name, record.S3.Object.URLDecodedKey, errGettingObj)
			return errGettingObj
		}

		users, errDecodingUsers := c.decodeUser(objectOutput.Body)
		if errDecodingUsers != nil {
			log.Fatalf("Error decoding object. Error: %s", errDecodingUsers.Error())
			return errDecodingUsers
		}

		return c.processUsers(ctx, users)
	}

	return nil
}

func (c *Controller) getS3Object(ctx context.Context, record events.S3EventRecord) (*s3.GetObjectOutput, error) {
	log.Printf("Getting S3 object from bucket %s and key %s", record.S3.Bucket.Name, record.S3.Object.URLDecodedKey)
	return c.s3.GetObject(ctx, &s3.GetObjectInput{
		Bucket: &record.S3.Bucket.Name,
		Key:    &record.S3.Object.URLDecodedKey,
	})
}

func (c *Controller) decodeUser(body io.ReadCloser) ([]User, error) {
	csvReader := csv.NewReader(body)
	csvReader.Comma = ';'
	decoder, errDec := csvutil.NewDecoder(csvReader)

	if errDec != nil {
		log.Print("Error decoding", errDec)
		return nil, errDec
	}

	var users []User
	for {
		var u User
		if err := decoder.Decode(&u); err == io.EOF {
			break
		}
		users = append(users, u)
	}
	return users, nil
}

func (c *Controller) processUsers(ctx context.Context, users []User) error {
	if len(users) > 0 {
		summaryUsers, errFiltering := filterUsers(users)
		if errFiltering != nil {
			log.Fatalf("Error parsing data %s", errFiltering.Error())
			return errFiltering
		}
		for _, v := range summaryUsers {
			errProcessing := c.srv.ProcessSummary(ctx, v)
			if errProcessing != nil {
				log.Fatalf("Error processing, error: ", errProcessing.Error())
				return errProcessing
			}
		}
	}
	return nil
}

func filterUsers(users []User) ([][]domain.User, error) {
	log.Println("Filtering users")
	var result [][]domain.User

	for i, v := range users {
		if i == 0 {
			sumUser, err := v.toUserSummary()
			if err != nil {
				log.Fatalf("Error converting user. Error: %s", err.Error())
				return nil, err
			}
			result = append(result, []domain.User{*sumUser})

		} else {
			sumUser, err := v.toUserSummary()
			if err != nil {
				log.Fatalf("Error converting user. Error: %s", err.Error())
				return nil, err
			}
			result = addUser(result, *sumUser)
		}

	}
	return result, nil
}

func addUser(result [][]domain.User, u domain.User) [][]domain.User {
	index := 0
	for i, _ := range result {
		if result[i][index].Id == u.Id {
			result[i] = append(result[i], u)
			return result
		}

		if i+1 == len(result) {
			return append(result, []domain.User{u})
		}

	}
	return result
}
