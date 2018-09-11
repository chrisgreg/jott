package models

type Editor struct {
	ID     uint `gorm:"unique";json:"-"`
	UserId uint `json:"-"`
	BlogId uint `json:"-"`
}
