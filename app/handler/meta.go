package handler

import (
	"github.com/jinzhu/gorm"
	"net/http"
)

func GetOverallStats(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	meta := struct {
		JottCount int
		BlogCount int
	}{}

	db.Table("jotts").Count(&meta.JottCount)
	db.Table("blogs").Count(&meta.BlogCount)
	respondJSON(w, http.StatusOK, meta)
}
