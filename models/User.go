package models

import (
	"fmt"
	"log"

	"github.com/DevitoDbug/golangAuthTemplate/utils"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Email        string `json:"email" validate:"required,email"`
	Password     string `json:"password" validate:"required,min=8"`
	SessionToken string `json:"sessionToken"`
	CSRFToken    string `json:"csrfToken"`
}

type ResponseUser struct {
	Email string `json:"email"`
}

var Storage = map[string]User{} // dummy storage

func (u *User) Store() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	if err != nil {
		errorContext := &utils.ErrorContext{
			Context: "User@Store",
			Value:   err.Error(),
		}

		log.Printf("Error while hashing password\n %s ", errorContext.Error())
		return err
	}

	Storage[u.Email] = User{
		Email:    u.Email,
		Password: string(hashedPassword),
	}

	fmt.Println("User stored successfully")
	log.Printf("User stored successfully")
	return nil
}

func (u *User) Show() ResponseUser {
	return ResponseUser{
		Email: u.Email,
	}
}
