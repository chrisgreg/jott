package models

import "time"

type Jott struct {
	ID      uint `gorm:"unique"`
	UserId  uint
	BlogId  uint
	Content string
	Created *time.Time
}