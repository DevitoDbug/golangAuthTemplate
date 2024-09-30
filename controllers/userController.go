package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/DevitoDbug/golangAuthTemplate/models"
	"github.com/DevitoDbug/golangAuthTemplate/utils"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		errorContext := &utils.ErrorContext{
			Context: "userController@CreateUser",
			Value:   err.Error(),
		}

		log.Printf("Error during validation\n %s ", errorContext.Error())
	}

	err = user.Store()
	if err != nil {
		errorContext := &utils.ErrorContext{
			Context: "userController@CreateUser",
			Value:   err.Error(),
		}

		log.Printf("Error while storing user\n %s ", errorContext.Error())
	}

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func ShowUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	for _, user := range models.Storage {
		users = append(users, user)
	}

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(users)
}
