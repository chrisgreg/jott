package handler

import (
	"encoding/json"
	"github.com/chrisgreg/jott/app/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func GetAllBlogsForUser(userId uint, db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	blogs := []models.Blog{}
	db.Where(&models.Blog{UserId: userId}).Find(&blogs)
	respondJSON(w, http.StatusOK, blogs)
}

func GetBlogByID(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idParam := vars["id"]
	blogID, err := strconv.ParseUint(idParam, 10, 32)

	if err != nil {
		respondError(w, http.StatusInternalServerError, "Couldn't get blog from id")
		return
	}

	blog := models.Blog{}
	db.LogMode(true)
	if err = db.
		Preload("User").
		Preload("Jotts").
		Preload("Jotts.User").
		Where(&models.Blog{ID: uint(blogID)}).
		Find(&blog).Error; err != nil {
		respondError(w, http.StatusNotFound, "Couldn't find blog with that id")
		return
	}

	blog.IncrementReadCount()
	db.Save(&blog)

	publicBlog := blog.ToPublicBlog()
	respondJSON(w, http.StatusOK, publicBlog)
}

func CreateNewBlog(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
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
		Title    string
		Subtitle string
		Private  bool
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

	// Create blog entry
	currentTime := time.Now()

	newBlog := models.Blog{
		UserId:    userId,
		Title:     JsonData.Title,
		Subtitle:  JsonData.Subtitle,
		Created:   &currentTime,
		Private:   JsonData.Private,
		ReadCount: 0,
	}

	if err = db.Create(&newBlog).Error; err != nil {
		respondError(w, http.StatusInternalServerError, "Couldn't create DB entry")
		return
	}

	publicBlog := newBlog.ToPublicBlog()

	respondJSON(w, http.StatusOK, publicBlog)
}
