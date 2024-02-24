package handlers

import (
	"log"
	"os"
)

type HandlerContext struct {
	Secret      []byte
}

func NewHandlerContext() (*HandlerContext) {
    // This is just an example where you want to store a secret in the HC.

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Fatal("Required constant 'JWT_SECRET' missing from .env file. Shutting down.")
	}

	return &HandlerContext{[]byte(secret)}
}
