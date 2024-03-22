package main

import (
	"log"
	"net/http"
	"submission3/handlers"
)

func main() {
	http.HandleFunc("/", handlers.StatusHandler)
	log.Println("Server started at localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
