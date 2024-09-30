package middleware

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/DevitoDbug/golangAuthTemplate/models"
	"github.com/DevitoDbug/golangAuthTemplate/utils"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var input struct {
			Email string `json:"email" ,validate:"required, email"`
		}

		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			errorContext := &utils.ErrorContext{
				Context: "Auth@authMiddleware",
				Value:   err.Error(),
			}

			log.Printf("Error during validation\n %s ", errorContext.Error())
			http.Error(w, "Validation error", http.StatusBadRequest)
		}

		user, exists := models.Storage[input.Email]
		if !exists {
			w.Header().Set("content-type", "application/json")
			http.Error(w, "User does not exist", http.StatusBadRequest)
			return
		}

		cookie, err := r.Cookie("session_token")
		if err != nil {
			http.Error(w, "Unauthorized: no session cookie found", http.StatusUnauthorized)
			return
		}

		sessionToken := cookie.Value
		if len(cookie.Value) == 0 || sessionToken != user.SessionToken {
			w.Header().Set("content-type", "application/json")
			http.Error(w, "Unauthorized: invalid session token", http.StatusUnauthorized)
			return
		}

		csrfToken := r.Header.Get("X-CSRF-Token")
		if len(csrfToken) == 0 || csrfToken != user.CSRFToken {
			w.Header().Set("content-type", "application/json")
			http.Error(w, "Unauthorized: invalid CSRF token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
