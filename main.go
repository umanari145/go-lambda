package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {

	if os.Getenv("IS_LAMBDA") != "" {
		lambda.Start(handle)
	} else {
		handle()
	}
}

func handle() (string, error) {
	fmt.Println(os.Getenv("TEST"))
	fmt.Println("test")
	//package 使えず・・エラーが出るわけでもなく何も出ない
	//sample.Test()
	log.Println("log だよ")
	return "hoge", nil
}
