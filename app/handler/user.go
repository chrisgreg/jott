package handler

import (
	"encoding/json"
	"fmt"
	"github.com/chrisgreg/jott/app/models"
	"github.com/chrisgreg/jott/app/utils"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetAllUsers(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	users := []models.User{}
	db.Find(&users)
	respondJSON(w, http.StatusOK, users)
}

func Login(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Couldn't read json body")
		return
	}

	var LoginDetails models.Login
	err = json.Unmarshal(body, &LoginDetails)

	if err != nil {
		respondError(w, http.StatusInternalServerError, "Couldn't read json body")
		return
	}

	// Get matching user from DB
	var FoundUser models.User
	db.Where(&models.User{Email: LoginDetails.Email}).First(&FoundUser)

	// Check passwords
	inputPassword := LoginDetails.Pass
	correctPassword := utils.CheckPasswordHash(inputPassword, FoundUser.Pass)

	if correctPassword == false {
		respondError(w, http.StatusForbidden, "Password incorrect")
		return
	}

	PublicUserDetails := FoundUser.GetPublicUser()
	tokenString, err := utils.CreateNewJWTToken(PublicUserDetails)

	if err != nil {
		respondError(w, http.StatusInternalServerError, "Account found but couldn't create JWT")
		return
	}

	respondJSON(w, http.StatusOK, utils.JWTResponse{Token: tokenString})
}

func CreateNewUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Couldn't read json body")
		return
	}

	var NewUser models.User
	err = json.Unmarshal(body, &NewUser)

	if err != nil {
		respondError(w, http.StatusInternalServerError, "Couldn't read json body")
		return
	}

	inputPassword := NewUser.Pass
	HashedPassword, err := utils.HashPassword(inputPassword)
	NewUser.Pass = HashedPassword

	if err != nil {
		respondError(w, http.StatusInternalServerError, "Couldn't read json body")
		return
	}

	err = db.Create(&NewUser).Error

	if err != nil && strings.Contains(err.Error(), "1062") {
		errorMessage := fmt.Sprintf("User with email: %s already exists", NewUser.Email)
		respondError(w, http.StatusInternalServerError, errorMessage)
		return
	} else if err != nil {
		respondError(w, http.StatusInternalServerError, "Couldn't create new user account")
		return
	}

	PublicUserDetails := NewUser.GetPublicUser()
	tokenString, err := utils.CreateNewJWTToken(PublicUserDetails)

	if err != nil {
		respondError(w, http.StatusInternalServerError, "Account created but couldn't create JWT")
		return
	}

	respondJSON(w, http.StatusOK, utils.JWTResponse{Token: tokenString})
}
