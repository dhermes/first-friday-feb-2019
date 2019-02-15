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
	publicKeyPEMBytes := []byte("-----BEGIN RSA PUBLIC KEY-----\nMIIBCgKCAQEAmBbzWGBpUqBlUKdt4tTe7ZWRzROY7gWt/qcsXP8eGZ+wnW/NEnYY\nCqtaDA+d/TL0zaHrzJEoy50dl0FqjyixUerK/90cWS8oOIghiYeMOovfoh4Zv7Tj\nnlzyfzQV9wn0xkOc7fGdapwZzDiabj3yunvFNhnvPJng6nJ0h7zUmvI+hIUcp5n4\nzH19kV/Hn1NVFzPDALQk+wHALlSBHEiOLAfh0EUl4FoKzztoe/FcBxIWYkEsutGS\nst1nSDQqEoOvykP05H1Ni7qHu7S7D5wIpztB9v+ag6blM/jSxP3M3AU3GHHd75pF\ns6jcRInHNGUj0NHSoaSrqoD/YIwEb0FZ4QIDAQAB\n-----END RSA PUBLIC KEY-----\n") // TODO: Get these from somewhere.
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
