package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"

	"github.com/dhermes/first-friday-feb-2019/pkg/verify"
)

type KeyPackage struct {
	KeyName string `json:"keyName"`
	Key     string `json:"key"`
}

type DbItem struct {
	UserId string     `json:"user_id"`
	Key    KeyPackage `json:"key"`
}

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (Response, error) {
	var err error

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-1")},
	)
	svc := dynamodb.New(sess)
	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("auth-keys"),
		Key: map[string]*dynamodb.AttributeValue{
			"user_id": {
				S: aws.String("8f055c6f-7e8d-4cb0-b243-65c7d633df36"), // Hardcoded; very wrong
			},
		},
	})
	item := DbItem{}
	_ = dynamodbattribute.UnmarshalMap(result.Item, &item)

	authorization := request.Headers["Authorization"]
	if !strings.HasPrefix(authorization, "Bearer ") {
		return Response{StatusCode: 401}, errors.New("Unauthorized.")
	}
	bearerTokenJWT := authorization[7:]
	publicKeyPEMBytes := []byte(item.Key.Key) // TODO: Get these from somewhere.
	var valid bool
	valid, err = verify.Verify(bearerTokenJWT, publicKeyPEMBytes, time.Now())
	if err != nil || !valid {
		return Response{StatusCode: 401}, errors.New("Invalid JWT.")
	}

	var buf bytes.Buffer
	body, err := json.Marshal(map[string]interface{}{
		"message": "who-am-i",
	})
	if err != nil {
		return Response{StatusCode: 404}, err
	}
	json.HTMLEscape(&buf, body)

	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type": "application/json",
			"X-Blend-Func": "who-am-i",
		},
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
