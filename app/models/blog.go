package models

import (
	"time"
)

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

type PublicBlog struct {
	ID        uint       `gorm:"unique"`
	User      PublicUser `json:"Author"`
	Title     string
	Subtitle  string
	Created   *time.Time
	Private   bool
	ReadCount uint
	Jotts     []PublicJott `gorm:"one2many:jotts"`
}

func (b *Blog) ToPublicBlog() PublicBlog {

	publicJotts := make([]PublicJott, len(b.Jotts))

	for i, value := range b.Jotts {
		publicJott := PublicJott{
			User:    value.User.GetPublicUser(),
			BlogId:  value.BlogId,
			Content: value.Content,
			Created: value.Created,
		}
		publicJotts[i] = publicJott
	}

	return PublicBlog{
		ID:        b.ID,
		User:      b.User.GetPublicUser(),
		Title:     b.Title,
		Subtitle:  b.Subtitle,
		Created:   b.Created,
		Private:   b.Private,
		ReadCount: b.ReadCount,
		Jotts:     publicJotts,
	}
}

func (b *Blog) IncrementReadCount() {
	b.ReadCount++
}
