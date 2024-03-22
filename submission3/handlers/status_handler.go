package handlers

import (
	_ "encoding/json"
	"html/template"
	"net/http"
	"submission3/updater"
)

type Status struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/status.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	status, err := updater.ReadStatusFromFile("status.json")
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
