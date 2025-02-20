package jsonutils

import (
	"encoding/json"
	"log"
	"os"
	"os/user"
)

type Session struct {
	Name string `json:"name"`
	Host string `json:"host"`
	Id   int    `json:"id"`
}

type Sessions struct {
	Sessions []Session `json:"sessions"`
}

func getSaveLocation() *string {
	usr, err := user.Current()

	if err != nil {
		log.Fatalf("Error getting current user")
	}

	homeDir := usr.HomeDir + "/.goSSH.json"

	return &homeDir
}

func Read() *Sessions {
	saveLocation := getSaveLocation()

	if saveLocation != nil {
		data, err := os.ReadFile(*saveLocation)

		if err != nil {
			initialSessions := Sessions{
				Sessions: []Session{{Name: "Test", Host: "Test", Id: 1}},
			}

			Write(initialSessions)
		}

		parsedSessions := Sessions{}
		json.Unmarshal(data, &parsedSessions)

		return &parsedSessions
	}

	return nil
}

func Write(sessions Sessions) {
	saveLocation := getSaveLocation()

	if saveLocation != nil {
		jsonData, err := json.MarshalIndent(sessions, "", " ")

		if err != nil {
			log.Fatalf("Failed to write file.")
		}
		os.WriteFile(*saveLocation, jsonData, 0644)
	}
}
