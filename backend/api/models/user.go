package models

import "net/http"

type User struct {
	Name      string `json:"name"`
	PhotoPath string `json:"photoPath"`
}

func (u *User) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
