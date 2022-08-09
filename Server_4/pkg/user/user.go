package user

import (
	"encoding/json"
	"errors"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/zarif007/GO-Serverless-CRUD/Server_4/pkg/validators"
)

var (
	ErrorfailedToFetchRecord = "Failed to fetch record"
	ErrorfailedToUnmarshalRecord = "Failed to Unmarshal record"
	ErrorInvalidUserData = "Invalid user data"
	ErrorInvalidEmail = "Invalid Email"
	ErrorCouldNotMarshalItem = "Could not marshal item"
	ErrorCouldNotDeleteItem = "Could not delete item"
	ErrorCouldNotDynamoPutItem = "Could not dynamo put item"
	ErrorUserAlreadyExists = "user.User already exists"
	ErrorUserDoesnotExists = "user.User Doesnot exists"
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
		return nil, errors.New(ErrorfailedToFetchRecord)
	}

	item := new(User)

	err = dynamodbattribute.UnmarshalMap(result.Item, item)

	if err != nil {
		return nil, errors.New((ErrorfailedToUnmarshalRecord))
	}

	return item, nil
}

func FetchUsers(tableName string, dynaclient dynamodbiface.DynamoDBAPI)(*[]User) {
	input := &dynamodb.ScanInput{
		tableName: aws.String(tableName)
	}

	result, err := dynaclient.Scan(input)

	if err != nil {
		return nil, errors.New((ErrorfailedToFetchRecord))
	}

	item := new([]user)

	err = dynamodbattribute.UnmarshalMap(result.Items, item)

	if err != nil {
		return nil, errors.New((ErrorfailedToUnmarshalRecord))
	}

	return item, nil
}

func CreateUser(req events.APIGatewayProxyRequest, tableName, dynaClient dynamodbiface.DynamoDBAPI)(*User, error) {
	var u User  

	json.Unmarshal([]byte(req.body), &u); err != nil {
		return nil, errors.New(ErrorInvalidUserData)
	}

	if !validators.IsEmailValid(u.Email) {
		return nil, errors.New(ErrorInvalidEmail)
	}

	currentUser, _ := FetchUser(u.Email, tableName, dynaClient)

	if currentUser != nil & len(currentUser.Email) != 0 {
		return nil, errors.New(ErrorUserAlreadyExists)
	}

	av, err := dynamodbattribute.MarshalMap(u)

	if err != nil {
		return nil, errors.New(ErrorCouldNotMarshalItem)
	}

	input := &dynamodb.PutItemInput{
		Item: av, 
		TableName: aws.String(tableName)
	}

	_, err := dynaClient.PutItem(input)

	if err != nil {
		return nil, errors.New(ErrorCouldNotDynamoPutItem)
	}
}

func UpdateUser() {

}

func DeleteUser() error {

}
