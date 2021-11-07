package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id uint `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password []byte `json:"-"`
	IsAdmin bool `json:"-"`
}

func (user *User) SetPassword(password string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	user.Password = hashedPassword
}