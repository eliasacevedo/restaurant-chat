package services

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"fastfoodrestaurant.com/api/models"
	"fastfoodrestaurant.com/api/models/responses"
)

type IBotService interface {
	AskQuestion(string) (*models.Message, error)
}

type BotService struct{}

const servicePath = "http://localhost:9091/v1/network/ask"
const jsonTypeRequest = "application/json; charset=utf-8"

func NewBotService() *BotService {
	return &BotService{}
}

func (b *BotService) AskQuestion(s string) (*models.Message, error) {
	message := &models.Message{
		Text: s,
	}

	messageRequest := &responses.MessageResponse{Message: message}

	jsonReq, err := json.Marshal(messageRequest)
	resp, err := http.Post(servicePath, jsonTypeRequest, bytes.NewBuffer(jsonReq))
	if err != nil {
		return &models.Message{
			Text: "El servicio no esta disponible",
		}, nil
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// bodyString := string(bodyBytes)

	var response responses.MessageResponse
	json.Unmarshal(bodyBytes, &response)
	return response.Message, nil
}
