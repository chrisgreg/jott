package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	ID    uint   `gorm:"unique"`
	Name  string `gorm:"unique" json:"name"`
	Email string `gorm:"unique"`
	Pass  string `json:'-'`
}

// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {
	// db.AutoMigrate(&User{})
	return db
}
