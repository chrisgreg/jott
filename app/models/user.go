package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	ID        uint   `gorm:"unique"`
	FirstName string `gorm:"unique" json:"first_name"`
	LastName  string `gorm:"unique" json:"last_name"`
	Username  string `gorm:"unique" json:"username"`
	Email     string `gorm:"unique" json:"email"`
	Pass      string `json:"pass"`
}

type PublicUser struct {
	FirstName string `gorm:"unique" json:"first_name"`
	LastName  string `gorm:"unique" json:"last_name"`
	Username  string `gorm:"unique" json:"username"`
}

type Login struct {
	Email string `json:"email"`
	Pass  string `json:"pass"`
}

// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {
	// db.AutoMigrate(&User{})
	return db
}

func (u User) GetPublicUser() PublicUser {
	return PublicUser{
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Username:  u.Username,
	}
}
