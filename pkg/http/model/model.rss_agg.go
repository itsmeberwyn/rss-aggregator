package model

import (
	"time"

	"github.com/google/uuid"
)

type UserModel struct {
	Id         uuid.UUID `json:"id"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	Name       string    `json:"name"`
	ApiKey     string    `json:"apikey"`
}

type FeedModel struct {
	Id              uuid.UUID `json:"id"`
	Created_at      time.Time `json:"created_at"`
	Updated_at      time.Time `json:"updated_at"`
	Url             string    `json:"url"`
	Name            string    `json:"name"`
	UserId          uuid.UUID `json:"user_id"`
	Last_fetched_at time.Time `json:"last_fetched_at"`
}

type RSSFeed struct {
	Link   string      `xml:"link"`
	Id     string      `xml:"id"`
	Title  string      `xml:"title"`
	Author []RSSAuthor `xml:"author"`
	Entry  []RSSItem   `xml:"entry"`
}

type RSSAuthor struct {
	Name string `xml:"name"`
	Uri  string `xml:"uri"`
}

type RSSItem struct {
	Id        string `xml:"id"`
	Title     string `xml:"title"`
	Link      Link   `xml:"link"`
	Published string `xml:"published"`
}

type Link struct {
	Key string `xml:"href,attr"`
}

type PostModel struct {
	Id           uuid.UUID `json:"id"`
	Created_at   time.Time `json:"created_at"`
	Updated_at   time.Time `json:"updated_at"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	Url          string    `json:"url"`
	Published_at time.Time `json:"published_at"`
	Feed_id      uuid.UUID `json:"feed_at"`
}
