package utils

import (
	"github.com/chrisgreg/jott/app/models"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"os"
)

var hmacSecret []byte

type JWTResponse struct {
	Token string
}

type UserClaims struct {
	models.PublicUser
	jwt.StandardClaims
}

func init() {
	envSecret := os.Getenv("hmacSecret")
	hmacSecret = []byte(envSecret)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CreateNewJWTToken(user models.PublicUser) (string, error) {
	claims := UserClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: 30000,
			Issuer:    "test",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(hmacSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
