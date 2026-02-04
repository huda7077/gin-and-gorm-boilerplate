package models

type AuthLogin struct {
	User
	Token string `json:"token"`
}
