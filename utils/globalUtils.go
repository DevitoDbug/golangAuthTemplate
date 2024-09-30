package utils

import (
	"crypto/rand"
	"encoding/base64"
	"log"
)

func GenerateToken(length int) string {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		errorContext := &ErrorContext{
			Context: "globalUtils@generateToken",
			Value:   err.Error(),
		}

		log.Fatalf("Failed to generate token\n %s ", errorContext)
	}
	return base64.URLEncoding.EncodeToString(bytes)
}
