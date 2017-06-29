package auth

import (
	"math/rand"

	"golang.org/x/crypto/bcrypt"
)

// GenerateRandomPassword is used to generate a random password based on
// letters from the letters variable and specified length
func GenerateRandomPassword(length int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// EncryptPassword is used to encrypt a password using bcrypt
func EncryptPassword(password string) string {
	pByte := []byte(password)
	hashedPassword, _ := bcrypt.GenerateFromPassword(pByte, 12)
	return string(hashedPassword)
}

// VerifyPassword is used to compare password against hash in db
func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
