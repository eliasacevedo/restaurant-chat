package models

type Message struct {
	Text           string   `json:"text"`
	Classification []string `json:"classification"`
}
