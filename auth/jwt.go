package auth

import (
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// GenerateJWT generates JWT token, based on roleID and userID
func GenerateJWT(userID uint, signingKey string, expiresIn time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(expiresIn).Unix(),
		"usr": userID,
	})
	return token.SignedString([]byte(signingKey))
}

// VerifyJWT is used to verify a jwt token, and return its claims
func VerifyJWT(theToken string, signingKey string) (claims map[string]interface{}, err error) {
	var token *jwt.Token

	// Try to parse key
	token, err = jwt.Parse(theToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(signingKey), nil
	})

	// If parsed but invalid, return error
	if err == nil && token.Valid {
		return nil, errors.New("Token is invalid")
	}

	return
}
