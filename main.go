package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {

	if os.Getenv("IS_LAMBDA") != "" {
		lambda.Start(handle)
	} else {
		req := loadSource()
		handle(req)
	}
}

//Req はリクエスト
type Req struct {
	//Message はメッセージ
	Messatge string `json:"message"`
	//Item は商品
	Items []Item `json:"items"`
}

//Item はItem
type Item struct {
	//Name は品名
	Name string `json:"name"`
	//Price は金額
	Price string `json:"price"`
}

func loadSource() Req {
	file, _ := os.Open("./sample_req.json")
	b, _ := ioutil.ReadAll(file)
	var sampleReq Req
	json.Unmarshal(b, &sampleReq)
	return sampleReq
}

func handle(sampleReq Req) (string, error) {

	fmt.Println(sampleReq)
	fmt.Println(os.Getenv("TEST"))
	fmt.Println("test")
	//lambdaだとpackage 使えず・・エラーが出るわけでもなく何も出ない
	//sample.Test()
	log.Println("log だよ")
	return "hoge", nil
}
