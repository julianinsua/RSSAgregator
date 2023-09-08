package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/julianinsua/RSSAgregator/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"apiKey"`
}

func (u *User) FromDB(dbUser database.User) *User {
	u.ID = dbUser.ID
	u.Name = dbUser.Name
	u.CreatedAt = dbUser.CreatedAt
	u.UpdatedAt = dbUser.UpdatedAt
	u.ApiKey = dbUser.ApiKey
	return u
}

func dbUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name:      dbUser.Name,
	}
}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserID    uuid.UUID `json:"userId"`
}

func (f *Feed) FromDB(dbFeed database.Feed) *Feed {
	f.ID = dbFeed.ID
	f.Name = dbFeed.Name
	f.Url = dbFeed.Url
	f.UserID = dbFeed.UserID
	f.CreatedAt = dbFeed.CreatedAt
	f.UpdatedAt = dbFeed.UpdatedAt

	return f
}

type Feeds []Feed

func (fs *Feeds) FromDB(dbFeeds []database.Feed) *Feeds {
	for _, dbFeed := range dbFeeds {
		fd := Feed{}
		fd.FromDB(dbFeed)
		*fs = append(*fs, fd)
	}
	return fs
}
