package responses

import (
	"fastfoodrestaurant/neuralnetwork/models"
	"net/http"
)

type MessageResponse struct {
	Message *models.Message `json:"message"`
}

func (m *MessageResponse) Bind(r *http.Request) error {
	return nil
}

func (m *MessageResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
