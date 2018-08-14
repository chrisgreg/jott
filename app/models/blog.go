package models

import "time"

type Blog struct {
	ID        uint `gorm:"unique"`
	UserId    uint
	Title     string
	Subtitle  string
	Created   *time.Time
	Private   bool
	ReadCount uint
}

func (b *Blog) IncrementReadCount() {
	b.ReadCount++
}
