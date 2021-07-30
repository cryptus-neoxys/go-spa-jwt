package routes

import (
	"github.com/cryptus-neoxys/go-spa-jwt/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	app.Post("/api/register", controllers.Register)
}