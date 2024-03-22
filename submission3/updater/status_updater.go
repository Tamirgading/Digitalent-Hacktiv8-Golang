package updater

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"time"
)

type Status struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

func UpdateStatusEvery15Seconds() {
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

func ReadStatusFromFile(filename string) (Status, error) {
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
