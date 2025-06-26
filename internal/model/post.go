package model

import (
	"time"

	"github.com/pradeepitm12/cb/bee/util"
)

type Post struct {
	ID              string
	Title           string
	Content         string
	Author          string
	PublicationDate time.Time
	LastModified    time.Time
	Tags            []string
}

func NewPost(title, content, author string, tags []string) *Post {
	return &Post{
		ID:              util.NewUUID(),
		Title:           title,
		Content:         content,
		Author:          author,
		PublicationDate: time.Now(),
		Tags:            tags,
	}
}
