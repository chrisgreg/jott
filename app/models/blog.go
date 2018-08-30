package models

import "time"

type Blog struct {
	ID        uint `gorm:"unique"`
	UserId    uint `json:"-"`
	User      User `json:"Author"`
	Title     string
	Subtitle  string
	Created   *time.Time
	Private   bool
	ReadCount uint
	Jotts     []Jott `gorm:"one2many:jotts"`
}

func (b *Blog) IncrementReadCount() {
	b.ReadCount++
}
