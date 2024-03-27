package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	_"path"
)
type Response struct{
	Success bool `json:"success"`
	Data []Person `json:"data"`
	Error string `json:"error"`
}
type Person struct{
	Id int `json:"id"`
	Name string `json:"name"`
	Address string `json:"address"`
	UUID string `json:"UUID"`
	Cards any `json:"Cards"`
	DeletedAt any `json:"DeletedAt"`
}

func main() {
	req, err := http.NewRequest("GET", "http://localhost:8082/person", nil)
	if err != nil {
		panic(err)
	}
	client := http.Client{}

	response, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	dataByte, err:=  io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	var r Response
	err = json.Unmarshal(dataByte, &r)
	if err != nil {
		panic(err)
	}

	fmt.Println(r.Data)
	p, ok := r.Data.([]any)
	if !ok {
		panic("not slice")
	}

	for _, v := range p {
		d, ok := v.(map[string]any)
		if !ok {
			panic("not map")
		}
		fmt.Println(d)
	}
	fmt.Println(p)
}