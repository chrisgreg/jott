package handler

import (
	"github.com/gorilla/context"
	"net/http"
	"strconv"

	"github.com/chrisgreg/jott/app/models"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
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

	respondJSON(w, http.StatusOK, blog)
}

func CreateNewBlog(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	val, ok := context.GetOk(r, "jwtClaims")
	if !ok {
		respondError(w, http.StatusInternalServerError, "Couldn't get claims")
		return
	}

}
