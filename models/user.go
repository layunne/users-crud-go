package models

import (
	"encoding/json"
)

type User struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (u *User) Bytes() []byte {

	bytes, err := json.Marshal(u)

	if err != nil {
		return nil
	}
	return bytes
}

func (u *User) String() string {

	return string(u.Bytes())
}

type CreateUser struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
