package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/luanpontes100/iceStorage/pkg/models"
)

func SaveTickets(t []models.Ticket) {
	if t != nil {
		file, err := os.Create("tickets.json")
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		defer file.Close()

		jsonEncoder := json.NewEncoder(file)
		err = jsonEncoder.Encode(t)
		if err != nil {
			fmt.Println("Error writing data to file:", err)
			return
		}

		fmt.Println("Data written to file.")
	}
}

func LoadTickets(t *[]models.Ticket) {
	// Read the JSON file into a byte slice
	jsonBytes, err := ioutil.ReadFile("tickets.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	err = json.Unmarshal(jsonBytes, t)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}
}
