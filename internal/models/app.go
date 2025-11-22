package models

import "time"

type User struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Surname      string    `json:"surname"`
	ThirdName    string    `json:"thirdName"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"passwordHash"`
	Birthdate    time.Time `json:"birthdate"`
}
