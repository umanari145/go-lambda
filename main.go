package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(handle)
}

func handle() (string, error) {
	fmt.Println("test")
	return "hoge", nil
}
