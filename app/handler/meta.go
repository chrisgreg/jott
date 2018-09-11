package handler

import (
	"github.com/chrisgreg/jott/app/models"
	"github.com/jinzhu/gorm"
	"net/http"
	"strings"
)

func GetOverallStats(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	meta := struct {
		JottCount  int `json:"jotts"`
		BlogCount  int `json:"blogs"`
		UserCount  int `json:"users"`
		TotalWords int `json:"total_words"`
	}{}

	var err error
	var Jotts []models.Jott
	totalWordCount := 0

	err = db.Table("jotts").Count(&meta.JottCount).Error
	err = db.Table("blogs").Count(&meta.BlogCount).Error
	err = db.Table("users").Count(&meta.UserCount).Error
	err = db.Find(&Jotts).Error

	if err != nil {
		respondError(w, http.StatusInternalServerError, "Couldn't fetch metastats")
		return
	}

	for _, value := range Jotts {
		totalWordCount += len(strings.Fields(value.Content))
	}

	meta.TotalWords = totalWordCount

	respondJSON(w, http.StatusOK, meta)
}
