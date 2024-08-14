//go:generate oapi-codegen --config=cfg.yaml ./api.yaml

package main

import (
	"log"

	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	middleware "github.com/oapi-codegen/echo-middleware"

	"github.com/robertjshirts/ratioed/backend/posts/api"
	"github.com/robertjshirts/ratioed/backend/posts/db"
	"github.com/robertjshirts/ratioed/backend/posts/utils"
)

func main() {
	db := db.NewDatabase()
	defer db.Close()

	server := api.NewServer(db)

	e := echo.New()

	swagger, err := api.GetSwagger()
	if err != nil {
		log.Fatal(err)
	}

	e.Use(middleware.OapiRequestValidator(swagger))

	api.RegisterHandlers(e, server)

	port := utils.GetEnv("PORT")
	address := "0.0.0.0:" + port
	log.Println("Listening on:", address)
	log.Fatal(e.Start(address))
}
