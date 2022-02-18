package utils

import (
	"encoding/json"
	"log"
	"time"
)

type ProviderMessage struct {
	Date    time.Time   `json:"date"`
	Action  string      `json:"action"`
	Message interface{} `json:"message"`
}

func (msg *ProviderMessage) GetJsonMessage() string {
	json, err := json.Marshal(msg)
	if err != nil {
		log.Fatal(err)
	}

	return string(json)
}
