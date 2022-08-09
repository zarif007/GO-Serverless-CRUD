package user

import (
	"encoding/json"
	"errors"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

var (
	ErrorfailedToFetchRecor = "Failed to fetch record"
)

type User struct {
	Email     string `json:"email"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

func FetchUser(email, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (*User, error) {

	input := &dynamodb.GetItemInput{
		Key: map[string]*dynadynamodb.AttributeValue{
			"email": {
				S: aws.String(email)
			}
		},
		TableName: aws.String(tableName)
	}

	result, err := dynaClient.GetItem(input)

	if err != nil {
		return nil, errors.New(ErrorfailedToFetchRecor)
	}

	item := new(User)

	err = dynamodbattribute.UnmarshalMap(result.Item, item)

	if err != nil {
		return nil, errors.New((ErrorfailedToFetchRecor))
	}

	return item, nil
}

func FetchUsers(tableName string, dynaclient dynamodbiface.DynamoDBAPI)(*[]User) {
	input := &dynamodb.ScanInput{
		tableName: aws.String(tableName)
	}

	dynaclient.Scan(input)
}

func CreateUser() {

}

func UpdateUser() {

}

func DeleteUser() error {

}
