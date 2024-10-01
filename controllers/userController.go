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
		http.Error(w, "Validation error", http.StatusBadRequest)
	}

	err = utils.Validate.Struct(user)
	if err != nil {
		errorContext := &utils.ErrorContext{
			Context: "userController@CreateUser",
			Value:   err.Error(),
		}

		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"errors": err.Error(),
		})

		log.Printf("Error during validation\n %s ", errorContext.Error())
		return
	}

	err = user.Store()
	if err != nil {
		errorContext := &utils.ErrorContext{
			Context: "userController@CreateUser",
			Value:   err.Error(),
		}

		log.Printf("Error while storing user\n %s ", errorContext.Error())
		http.Error(w, "Error while storing user", http.StatusInternalServerError)
	}

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func ShowUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.ResponseUser
	for _, user := range models.Storage {
		users = append(users, user.Show())
	}

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(users)
}
