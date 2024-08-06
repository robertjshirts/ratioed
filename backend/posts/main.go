package main

import (
	_ "github.com/joho/godotenv/autoload"

	"github.com/robertjshirts/ratioed/backend/posts/db"
)

func main() {
	db := db.NewPostsDb()
	defer db.Close()
}
