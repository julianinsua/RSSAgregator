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

type FeedFollow struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	UserID    uuid.UUID `json:"userId"`
	FeedID    uuid.UUID `json:"feedId"`
}

func (ff *FeedFollow) FromDB(dbFeedFollow database.FeedFollow) *FeedFollow {
	ff.ID = dbFeedFollow.ID
	ff.CreatedAt = dbFeedFollow.CreatedAt
	ff.UpdatedAt = dbFeedFollow.UpdatedAt
	ff.UserID = dbFeedFollow.FeedID
	ff.UserID = dbFeedFollow.UserID
	return ff
}

type FeedFollows []FeedFollow

func (ffs *FeedFollows) FromDB(dbFeedFollows []database.FeedFollow) *FeedFollows {
	for _, dbFeedFollow := range dbFeedFollows {
		ff := FeedFollow{}
		ff.FromDB(dbFeedFollow)
		*ffs = append(*ffs, ff)
	}
	return ffs
}

type Post struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	PublishedAt time.Time `json:"publishedAt"`
	Url         string    `json:"url"`
	FeedID      uuid.UUID `json:"feedId"`
}

func (p *Post) FromDB(dbPost database.Post) *Post {
	p.ID = dbPost.ID
	p.CreatedAt = dbPost.CreatedAt
	p.UpdatedAt = dbPost.UpdatedAt
	p.Title = dbPost.Title
	p.Description = dbPost.Description.String
	p.PublishedAt = dbPost.PublishedAt
	p.Url = dbPost.Url
	p.FeedID = dbPost.FeedID
	return p
}

type Posts []Post

func (ps *Posts) FromDB(dbPosts []database.Post) *Posts {
	for _, dbPost := range dbPosts {
		p := Post{}
		p.FromDB(dbPost)

		*ps = append(*ps, p)
	}
	return ps
}
