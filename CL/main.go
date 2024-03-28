package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/sony/gobreaker"
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
	for {
		time.Sleep(1 * time.Second)
		result, err := get()
		log.Println(result, err)
	}
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

func get() ([]byte, error){
	// cbs = circuit breaker setting
	cbs := gobreaker.Settings{
		Name: "cl",
		Interval: 5* time.Second,
		Timeout: 7* time.Second,
		ReadyToTrip: func(counts gobreaker.Counts) bool{
			failurRatio := float64(counts.TotalFailures) / float64(counts.Requests)
			return counts.Requests >= 3 && failurRatio >= 0.6
		},
		OnStateChange: func(_ string, from, to gobreaker.State) {
			log.Println("state changed from", from.String(), "to", to.String())
		},
	}
	cb := gobreaker.NewCircuitBreaker(cbs)

	dataByte, err := cb.Execute(func() (interface{}, error) {
		response, err := http.Get("http://localhost:8082/hello")
		if err != nil {
			return nil, err
		}
		defer response.Body.Close()
	
		log.Println(response.Status)
	
		dataByte, err := io.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}
	
		fmt.Println(string(dataByte))
		if response.StatusCode == http.StatusInternalServerError{
			return nil, errors.New("internal server error")
		}
		return dataByte, nil
	})

	if err != nil {
		return nil, err 
	}

	return dataByte.([]byte), nil
	//log.Println(body, err)
}