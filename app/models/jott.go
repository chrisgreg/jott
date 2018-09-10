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

type PublicJott struct {
	User    PublicUser
	BlogId  uint
	Content string
	Created *time.Time
}

func (j *Jott) ToPublicJott() PublicJott {
	return PublicJott{
		User:    j.User.GetPublicUser(),
		BlogId:  j.BlogId,
		Content: j.Content,
		Created: j.Created,
	}
}
