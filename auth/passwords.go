package auth

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"time"

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

// GenerateConfirmHash generates a hash that can be used in for
// example password confirm mails
func GenerateConfirmHash() (string, error) {
	rData := make([]byte, 128)
	if _, err := rand.New(rand.NewSource(int64(time.Now().Nanosecond()))).Read(rData); err != nil {
		return "", err
	}
	hash := md5.Sum(rData)

	return hex.EncodeToString(hash[:]), nil
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
