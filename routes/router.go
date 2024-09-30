package router

import (
	"net/http"

	"github.com/DevitoDbug/golangAuthTemplate/controllers"
	"github.com/DevitoDbug/golangAuthTemplate/middleware"
)

func Router(w http.ResponseWriter, r *http.Request) {
	method := r.Method

	switch r.URL.Path {
	case "/register":
		if method == "POST" {
			controllers.CreateUser(w, r)
		} else {
			http.Error(w, "Bad request", http.StatusBadRequest)
		}

	case "/login":
		if method == "POST" {
			controllers.Login(w, r)
		} else {
			http.Error(w, "Bad request", http.StatusBadRequest)
		}

	case "/logout":
		if method == "POST" {
			controllers.LogOut(w, r)
		} else {
			http.Error(w, "Bad request", http.StatusBadRequest)
		}

	case "/show-users":
		if method == "GET" {
			middleware.Auth(http.HandlerFunc(controllers.ShowUsers)).ServeHTTP(w, r)
		} else {
			http.Error(w, "Bad request", http.StatusBadRequest)
		}

	}
}
