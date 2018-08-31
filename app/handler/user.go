package handler

import (
	"encoding/json"
	"fmt"
	"github.com/chrisgreg/jott/app/models"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetAllUsers(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	users := []models.User{}
	db.Find(&users)
	respondJSON(w, http.StatusOK, users)
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

	HashedPassword, _ := HashPassword(NewUser.Pass)
	NewUser.Pass = HashedPassword

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
	respondJSON(w, http.StatusOK, PublicUserDetails)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
