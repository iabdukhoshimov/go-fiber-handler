package main

import (
	"log"

	"github.com/go-fiber/config"
	"github.com/gofiber/fiber/v2"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	app := fiber.New()
	// db := db.Init(c.DBUrl)

	// books.RegisterRoutes(app, db)

	app.Listen(c.Port)
}
