package handler

import (
	"fmt"
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
	db.
		Joins("INNER JOIN jotts on blogs.id=blogJotts.blog_id").
		Where(&models.Blog{ID: uint(blogID)}).
		Find(&blog)

	fmt.Println(blog)

	blog.IncrementReadCount()
	db.Save(&blog)

	respondJSON(w, http.StatusOK, blog)
}
