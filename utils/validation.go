package utils

import "github.com/go-playground/validator"

var Validate *validator.Validate

func init() {
	// Docs recommend having only one instance of Validate
	Validate = validator.New()
}
