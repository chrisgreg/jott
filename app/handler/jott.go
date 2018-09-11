package handler

import (
	"encoding/json"
	"github.com/chrisgreg/jott/app/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"net/http"
	"time"
)

func CreateJott(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	// Grab claims from context
	val, ok := context.GetOk(r, "jwtClaims")
	if !ok {
		respondError(w, http.StatusInternalServerError, "Couldn't get claims")
		return
	}

	// Parse data from post
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Couldn't read json body")
		return
	}

	JsonData := struct {
		BlogId  uint
		Content string
	}{}

	err = json.Unmarshal(body, &JsonData)

	if err != nil {
		respondError(w, http.StatusInternalServerError, "Couldn't read json body")
		return
	}

	// Get username from claims
	claimsValues := val.(jwt.MapClaims)
	username := claimsValues["username"].(string)

	// Get user id from username
	var userId uint
	row := db.Table("users").Where("username = ?", username).Select("id").Row()
	row.Scan(&userId)

	// Check if the user is an editor of the blog
	var editor models.Editor
	var count uint
	db.Where("blog_id = ? AND user_id >= ?", JsonData.BlogId, userId).First(&editor).Count(&count)
	if count <= 0 {
		respondError(w, http.StatusForbidden, "Don't have permission to add to this blog")
		return
	}

	// Create jott entry
	currentTime := time.Now()

	jott := models.Jott{
		UserId:  userId,
		BlogId:  JsonData.BlogId,
		Content: JsonData.Content,
		Created: &currentTime,
	}

	if err = db.Create(&jott).Error; err != nil {
		respondError(w, http.StatusInternalServerError, "Couldn't create DB entry")
		return
	}

	publicJott := jott.ToPublicJott()
	respondJSON(w, http.StatusOK, publicJott)
}
