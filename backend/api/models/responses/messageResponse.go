package responses

import (
	"net/http"

	"fastfoodrestaurant.com/api/models"
)

type MessageResponse struct {
	Message *models.Message `json:"message"`
}

func (m *MessageResponse) Render(rw http.ResponseWriter, r *http.Request) error {
	return nil
}

func (u *MessageResponse) Bind(r *http.Request) error {
	return nil
}
