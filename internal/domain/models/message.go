package models

import (
	"encoding/json"
	"errors"
	"time"
)

type MessageData struct {
	Pattern string  `json:"pattern"`
	Message Message `json:"data"`
}

type Message struct {
	Identifier  string `json:"identifier"`
	Message     string `json:"message"`
	PhoneNumber string `json:"phoneNumber"`
	CampaignId  string `json:"campaignId"`
	Total       int    `json:"total"`
}

func CreateMessage(data []byte) (*Message, error) {
	var message MessageData
	err := json.Unmarshal(data, &message)
	if err != nil {
		return nil, errors.New("Failed unmarshalling message: " + err.Error())
	}

	return &message.Message, nil
}

func (e *Message) Prepare() map[string]interface{} {
	now := time.Now().UTC()

	message := map[string]interface{}{
		"identifier":   e.Identifier,
		"message":      e.Message,
		"phone_number": e.PhoneNumber,
		"campaign_id":  e.CampaignId,
		"created_at":   now,
		"updated_at":   now,
		"deleted":      false,
	}

	return message
}
