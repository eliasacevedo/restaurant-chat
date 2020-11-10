package models

import "net/http"

type Message struct {
	Text           string   `json:"text"`
	Classification []string `json:"classification"`
}

func (m *Message) Render(rw http.ResponseWriter, r *http.Request) error {
	return nil
}

func (u *Message) Bind(r *http.Request) error {
	return nil
}
