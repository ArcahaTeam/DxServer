package services

import (
	"crypto/sha256"
	"fmt"
)

func Login(username, password string) (bool, string) {
	if username == "admin" && password == "password" {
		token := generateToken(username)
		return true, token
	}
	return false, ""
}

func generateToken(username string) string {
	hash := sha256.Sum256([]byte(username + "secret_key"))
	return fmt.Sprintf("%x", hash[:])
}
