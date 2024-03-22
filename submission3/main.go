package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type Status struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

func main() {
	go updateStatusEvery15Seconds()

	http.HandleFunc("/", statusHandler)
	fmt.Println("Server started at localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func updateStatusEvery15Seconds() {
	for {
		status := Status{
			Water: rand.Intn(100) + 1,
			Wind:  rand.Intn(100) + 1,
		}

		file, err := os.Create("status.json")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		encoder := json.NewEncoder(file)
		err = encoder.Encode(status)
		if err != nil {
			log.Fatal(err)
		}

		time.Sleep(15 * time.Second)
	}
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/status.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	status, err := readStatusFromFile("status.json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func readStatusFromFile(filename string) (Status, error) {
	var status Status

	file, err := os.Open(filename)
	if err != nil {
		return status, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&status)
	if err != nil {
		return status, err
	}

	return status, nil
}
