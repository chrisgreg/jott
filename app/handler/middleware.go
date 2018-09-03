package handler

import (
	"fmt"
	"github.com/chrisgreg/jott/app/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"log"
	"net/http"
	"strings"
)

func Log(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		f(w, r)
	}
}

func Protected(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authToken := r.Header.Get("Authorization")
		authArr := strings.Split(authToken, " ")

		if len(authArr) != 2 {
			log.Println("Authentication header is invalid: " + authToken)
			respondError(w, http.StatusUnauthorized, "Authentication header is invalid")
			return
		}

		jwtToken := authArr[1]

		token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return utils.HMACSecret, nil
		})

		if err != nil {
			log.Println("Couldn't parse token: " + authToken)
			respondError(w, http.StatusUnauthorized, "Couldn't parse token")
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			fmt.Println(claims)
			context.Set(r, "jwtClaims", claims)
			f(w, r)
		} else {
			respondError(w, http.StatusUnauthorized, "Token is not valid")
			return
		}
	}
}
