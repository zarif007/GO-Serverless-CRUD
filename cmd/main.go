package main

import (
	"os"

	"github.com/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/lambda"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {
	region := os.Getenv("AWS_REGION")
	awsSession, err := session.newSession(&aws.Config{
		Region: aws.String(region)})

	if err != nil {
		return
	}

	dynoClient = dynamodb.New(awsSession)
	lambda.Start(handler)

}
