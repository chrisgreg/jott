package models

import (
	"math"
	"strings"
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
	ID         uint       `gorm:"unique"`
	User       PublicUser `json:"Author"`
	Title      string
	Subtitle   string
	Created    *time.Time
	Private    bool
	ReadCount  uint
	Jotts      []PublicJott `gorm:"one2many:jotts"`
	TotalWords int
	TimeToRead int
}

func (b *Blog) ToPublicBlog() PublicBlog {

	totalWords := 0

	publicJotts := make([]PublicJott, len(b.Jotts))
	for i, value := range b.Jotts {
		publicJott := PublicJott{
			User:    value.User.GetPublicUser(),
			BlogId:  value.BlogId,
			Content: value.Content,
			Created: value.Created,
		}
		totalWords += len(strings.Fields(value.Content))
		publicJotts[i] = publicJott
	}

	timeToRead := calculateTimeToRead(totalWords)

	return PublicBlog{
		ID:         b.ID,
		User:       b.User.GetPublicUser(),
		Title:      b.Title,
		Subtitle:   b.Subtitle,
		Created:    b.Created,
		Private:    b.Private,
		ReadCount:  b.ReadCount,
		Jotts:      publicJotts,
		TotalWords: totalWords,
		TimeToRead: timeToRead,
	}
}

func (b *Blog) IncrementReadCount() {
	b.ReadCount++
}

func calculateTimeToRead(words int) int {
	const avgWordsPerMinute = 200
	minutes := float64(words) / avgWordsPerMinute
	minutesToRead := math.Ceil(minutes)
	return int(minutesToRead)
}
