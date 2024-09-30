package main

import (
	"log"
	"net/http"

	router "github.com/DevitoDbug/golangAuthTemplate/routes"
	"github.com/DevitoDbug/golangAuthTemplate/utils"
	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

func main() {
	Validate = validator.New(validator.WithRequiredStructEnabled())

	port := ":8080"
	var appRouter http.HandlerFunc = router.Router
	log.Printf("Starting server in port %s...", port)
	err := http.ListenAndServe(port, appRouter)
	if err != nil {
		errorContext := &utils.ErrorContext{
			Context: "main",
			Value:   err.Error(),
		}

		log.Printf("Error listening to the port\n %s ", errorContext.Error())
	}

}
