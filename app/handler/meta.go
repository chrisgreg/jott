package handler

import (
	"github.com/jinzhu/gorm"
	"net/http"
)

func GetOverallStats(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	meta := struct {
		JottCount int
		BlogCount int
		UserCount int
	}{}

	var err error
	err = db.Table("jotts").Count(&meta.JottCount).Error
	err = db.Table("blogs").Count(&meta.BlogCount).Error
	err = db.Table("users").Count(&meta.UserCount).Error

	if err != nil {
		respondError(w, http.StatusInternalServerError, "Couldn't fetch metastats")
		return
	}

	respondJSON(w, http.StatusOK, meta)
}
