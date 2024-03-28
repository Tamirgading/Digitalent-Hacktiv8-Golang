package main

import (
	"fmt"
	"log"
	"net/http"
)

var isSuccess bool = true

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		log.Println("receive request")
		if isSuccess {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("success response"))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("fail response"))
		}	
	})
	http.HandleFunc("/toggle", func(w http.ResponseWriter, r *http.Request) {
		if isSuccess {
			fmt.Println("success change to false")
			isSuccess = false
		} else {
			fmt.Println("success change to true")
			isSuccess = true
		}
	})
	http.ListenAndServe("localhost:8082", nil)
}