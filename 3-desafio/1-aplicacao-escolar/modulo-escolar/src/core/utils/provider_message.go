package utils

import (
	"encoding/json"
	"log"
)

type ProviderMessage struct {
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
