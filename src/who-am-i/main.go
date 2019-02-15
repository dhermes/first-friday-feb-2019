package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/dhermes/first-friday-feb-2019/pkg/verify"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (Response, error) {
	var err error

	authorization := request.Headers["Authorization"]
	if !strings.HasPrefix(authorization, "Bearer ") {
		return Response{StatusCode: 401}, errors.New("Unauthorized.")
	}
	bearerTokenJWT := authorization[7:]
	publicKeyPEMBytes := []byte("-----BEGIN RSA PUBLIC KEY-----\nMIIBCgKCAQEApbRv8NhJ8jZ2fHK8FlomklkYCb1jxbLTSjNQ8IUdCJ1TBaP0u7sh\n1rvyKhK0TPwx5tZkm4ZtgACKmw8Wfok8Lf8OkYPOYdZ1Lj9ftxIS+B8/S/tld73x\nqwRj3S+iKUH8UNKtVovgdUsojlBvuMe5fwRw/QL0/cO3iwl73vMFjn6MPQWUeZzO\n9S8peDf/HogVhHcO2k2wsunfepgX0cbZPfwwtOQ9ZJq1+RcUNVQaR5EU6CoTnc6l\ntaesywuZi4a4OaN/eMD+ZXX6JEldq3t/PZP2tA3tLtWtkVHJhN6pb8vhzeEEftqV\n5FtrGzJM+H9lrFNJmfNAL05Rc14GZZrqiQIDAQAB\n-----END RSA PUBLIC KEY-----\n") // TODO: Get these from somewhere.
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
