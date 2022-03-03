package messaging

import (
	"encoding/json"
	"log"
)

type ProviderMessage struct {
	Action  string `json:"action"`
	Message string `json:"message"`
}

func NewProviderMessage(action string, message string) *ProviderMessage {
	return &ProviderMessage{Action: action, Message: message}
}

func (msg *ProviderMessage) GetJson() string {
	message, err := json.Marshal(msg)
	if err != nil {
		log.Fatal(err)
	}

	return string(message)
}
