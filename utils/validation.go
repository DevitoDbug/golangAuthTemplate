package utils

import "github.com/go-playground/validator"

var Validate *validator.Validate

func Init() {
	// Docs recommend having only one instance of Validate
	Validate = validator.New()
}
