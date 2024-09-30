package models

import (
	"fmt"
	"log"

	"github.com/DevitoDbug/golangAuthTemplate/utils"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Email        string `json:"email" validate:"required, email"`
	Password     []byte `json:"password" validate:"required, len=8"`
	SessionToken string `json:"sessionToken"`
	CSRFToken    string `json:"csrfToken"`
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
		Password: hashedPassword,
	}

	fmt.Println("User stored successfully")
	log.Printf("User stored successfully")
	return nil
}
