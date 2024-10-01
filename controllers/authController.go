package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/DevitoDbug/golangAuthTemplate/models"
	"github.com/DevitoDbug/golangAuthTemplate/utils"
	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var userInput models.User
	err := json.NewDecoder(r.Body).Decode(&userInput)
	if err != nil {
		errorContext := &utils.ErrorContext{
			Context: "userController@CreateUser",
			Value:   err.Error(),
		}

		log.Printf("Error during validation\n %s ", errorContext.Error())
	}

	err = utils.Validate.Struct(userInput)
	if err != nil {
		errorContext := &utils.ErrorContext{
			Context: "authController@Login",
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

	// Check if the user exists in the Storage by email
	storedUser, exists := models.Storage[userInput.Email]
	if !exists {
		log.Printf("User not found")
		http.Error(w, "User does not exist", http.StatusUnauthorized)
		return
	}

	// Compare the hashed password with the provided password
	err = bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(userInput.Password))
	if err != nil {
		log.Printf("Invalid password")
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	sessionToken := utils.GenerateToken(32)
	csrfToken := utils.GenerateToken(32)
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    sessionToken,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})
	http.SetCookie(w, &http.Cookie{
		Name:     "csrfToken_token",
		Value:    csrfToken,
		Expires:  time.Now().Add(5 * time.Minute),
		HttpOnly: false,
	})
	storedUser.SessionToken = sessionToken
	storedUser.CSRFToken = csrfToken
	models.Storage[userInput.Email] = storedUser

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(storedUser)
}

func LogOut(w http.ResponseWriter, r *http.Request) {
	var userInput struct {
		Email string `json:"email" ,validate:"required, email"`
	}

	err := json.NewDecoder(r.Body).Decode(&userInput)
	if err != nil {
		errorContext := &utils.ErrorContext{
			Context: "authController@LogOut",
			Value:   err.Error(),
		}

		log.Printf("Error during validation\n %s ", errorContext.Error())
		http.Error(w, "Validation error", http.StatusBadRequest)
	}

	// Validate the user struct
	err = utils.Validate.Struct(userInput)
	if err != nil {
		errorContext := &utils.ErrorContext{
			Context: "authController@LogOut",
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

	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "csrfToken_token",
		Value:    "",
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: false,
	})

	// Delete the tokens from storage
	user, exists := models.Storage[userInput.Email]
	if !exists {
		w.Header().Set("content-type", "application/json")
		http.Error(w, "User does not exist", http.StatusBadRequest)
	}

	user.CSRFToken = ""
	user.SessionToken = ""
	models.Storage[userInput.Email] = user

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "User logged out successfully"})
}
