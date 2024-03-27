package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Response struct {
	Success bool     `json:"success"`
	Data    []Person `json:"data"`
	Error   string   `json:"error"`
}

type Person struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Address   string `json:"address"`
	UUID      string `json:"UUID"`
	Cards     []CreditCard `json:"Cards"`
	DeletedAt interface{} `json:"DeletedAt"`
}

type CreditCard struct {
	CardNumber string `json:"cardNumber"`
	PersonUUID string `json:"personUUID"`
}

func main() {
	//post()
	get()
}

func post(){
	var data string = `
	{
		"name": "tegar",
		"address": "lampung"
	}`
	req, err := http.NewRequest("POST", "http://localhost:8082/person", bytes.NewBuffer([]byte(data)))
	if err != nil {
		panic(err)
	}
	client := http.Client{}

	response, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	resultByte, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(resultByte))
}

func get(){
	req, err := http.NewRequest("GET", "http://localhost:8082/person", nil)
	if err != nil {
		panic(err)
	}
	client := http.Client{
		Timeout: 1 * time.Second,
	}

	response, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	dataByte, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	var r Response
	err = json.Unmarshal(dataByte, &r)
	if err != nil {
		panic(err)
	}

	fmt.Println(r)
}