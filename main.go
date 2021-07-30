package main

import (
	"github.com/cryptus-neoxys/go-spa-jwt/database"
	"github.com/cryptus-neoxys/go-spa-jwt/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {

	database.Connect()

	app := fiber.New()

	routes.Setup(app)

	app.Listen(":8080")
}
