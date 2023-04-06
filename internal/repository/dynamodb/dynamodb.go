package dynamodb

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/summary/internal/core/domain"
	"github.com/summary/internal/core/port/outbound"
	"log"
)

type repositoryService struct {
	dbClient  dynamodb.Client
	tableName string
}

func NewRepositoryService(db dynamodb.Client, tableName string) outbound.RepositoryService {
	return &repositoryService{
		dbClient:  db,
		tableName: tableName,
	}
}

func (r *repositoryService) SaveItem(ctx context.Context, user domain.BalanceUser) error {
	userRepo := toRepositoryUser(user)
	dynamoItem, errMarshal := attributevalue.MarshalMap(userRepo)
	if errMarshal != nil {
		log.Fatalf("Error Marshaling user. Error: %s", errMarshal.Error())
		return errMarshal
	}
	_, errSaving := r.dbClient.PutItem(ctx, &dynamodb.PutItemInput{
		Item:      dynamoItem,
		TableName: &r.tableName,
	})
	return errSaving
}

func toRepositoryUser(user domain.BalanceUser) RepositoryUser {
	return RepositoryUser{
		Email:        user.EmailUser,
		TotalBalance: user.TotalBalance,
		CreditAvg:    user.TotalCreditBalance,
		DebitAvg:     user.TotalDebitBalance,
	}
}
