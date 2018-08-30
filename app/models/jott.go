package models

import "time"

type Jott struct {
	ID      uint `gorm:"unique";json:"-"`
	User    User
	UserId  uint `json:"-"`
	BlogId  uint `json:"-"`
	Content string
	Created *time.Time
}
