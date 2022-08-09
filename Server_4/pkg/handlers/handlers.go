package handlers

import (
	"net/http"

	"github.com/aws-sdk-go/aws"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/zarif007/GO-Serverless-CRUD/Server_4/handlers"
	"github.com/zarif007/GO-Serverless-CRUD/Server_4/pkg/user"
)

var ErrorMethodNotAllowed = "Method not allowed"

type ErrorBody struct {
	ErrorMsg *string `json:"error,omitempty"`
}

func GetUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI)
	(*events.APIGatewayProxyResponse, error) 
	{
		email := req.QueryStringParamters["email"]

		if len(email) > 0 {
			result, err := user.FetchUser(email, tableName, dynaClient)

			if err != nil {
				return apiResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
			}
			
			return apiResponse(http.StatusOK, result)
		}
		
		result, err := user.FetchUsers(tableName, dynaClient)

		if err != nil {
			return apiResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
		}

		return apiResponse(http.StatusOK, result)

}

func CreateUser() {

}

func UpdateUser() {

}

func DeleteUser() {

}

func UnhanledMethod()(*events.APIGatewayProxyResponse, error) {

}
