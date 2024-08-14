package db

import (
	_ "github.com/lib/pq"

	"github.com/robertjshirts/ratioed/backend/posts/api"
)

type tempPost struct {
	api.Post
	api.Content
	AccountId int `db:"account_id"`
}

func (p *tempPost) ToPost() *api.Post {
	post := p.Post
	post.Content = p.Content
	return &post
}
