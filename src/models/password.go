package models

type Password struct {
	NewPassword string `json:"newpassword"`
	Password    string `json:"password"`
}
