//go:generate oapi-codegen --config=cfg.yaml ./api.yaml

package main

import (
	"log"

	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"

	"github.com/robertjshirts/ratioed/backend/posts/api"
	"github.com/robertjshirts/ratioed/backend/posts/db"
)

func main() {
	db := db.NewDatabase()
	defer db.Close()

	server := api.NewServer(db)

	e := echo.New()

	api.RegisterHandlers(e, server)

	log.Println("Listening on 0.0.0.0:8080")
	log.Fatal(e.Start("0.0.0.0:8080"))
}
