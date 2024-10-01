package main

import (
	"log"
	"net/http"

	router "github.com/DevitoDbug/golangAuthTemplate/routes"
	"github.com/DevitoDbug/golangAuthTemplate/utils"
)

func main() {
	utils.Init()

	port := ":8080"
	var appRouter http.HandlerFunc = router.Router
	log.Printf("Starting server in port %s...", port)
	err := http.ListenAndServe(port, appRouter)
	if err != nil {
		errorContext := &utils.ErrorContext{
			Context: "main",
			Value:   err.Error(),
		}

		log.Fatalf("Error listening to the port\n %s ", errorContext.Error())
	}
}
