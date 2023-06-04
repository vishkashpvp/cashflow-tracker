package models

type User struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Provider string `json:"provider"`
	Picture  string `json:"picture"`
}
