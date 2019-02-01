package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"

	"bytes"
	"context"
	"encoding/asn1"
	"encoding/json"
	"encoding/pem"

	uuid "github.com/satori/go.uuid"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context) (Response, error) {
	// private key will be written to "keyName", public key will be written to "keyName.pub"
	keyName := uuid.Must(uuid.NewV4()).String()

	publicKey, privateKey := createKeys(keyName)
	savePubKey(publicKey, keyName)
	var buf bytes.Buffer

	body, err := json.Marshal(map[string]interface{}{
		"publicKey":  publicKey,
		"privateKey": privateKey,
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
			"X-Blend-Func": "gimme-key",
		},
	}

	return resp, nil
}

func createKeys(keyName string) (KeyPackage, KeyPackage) {
	// Generate RSA key pair
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	checkError(err)

	pubPackage := packagePublicKey(keyName, key.PublicKey)
	privatePackage := packagePrivateKey(keyName, key)
	return pubPackage, privatePackage
}

func packagePublicKey(keyName string, pubKey rsa.PublicKey) KeyPackage {
	asn1Bytes, err := asn1.Marshal(pubKey)
	var pemKey = &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: asn1Bytes,
	}
	checkError(err)
	return KeyPackage{
		KeyName: keyName + ".pub",
		Key:     string(pem.EncodeToMemory(pemKey)),
	}
}

func packagePrivateKey(keyName string, privateKey *rsa.PrivateKey) KeyPackage {
	var pemKey = &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}
	return KeyPackage{
		KeyName: keyName,
		Key:     string(pem.EncodeToMemory(pemKey)),
	}
}

func savePubKey(publicKeyPackage KeyPackage, userId string) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)
	checkError(err)
	// Create DynamoDB client
	svc := dynamodb.New(sess)

	newDbItem := DbItem{
		UserId: userId,
		Key:    publicKeyPackage,
	}

	av, err := dynamodbattribute.MarshalMap(newDbItem)
	checkError(err)

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("auth-keys"),
	}

	_, err = svc.PutItem(input)

	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

type KeyPackage struct {
	KeyName string `json:"keyName"`
	Key     string `json:"key"`
}

type DbItem struct {
	UserId string     `json:"user_id"`
	Key    KeyPackage `json:"key"`
}

func main() {
	lambda.Start(Handler)
}
