package Routers

import (
	"Auth/Controllers"
	"github.com/gofiber/fiber/v2"
)

func RoutingUser() {
	app := fiber.New()

	app.Get("/users", Controllers.GetAllUsers)
	app.Post("/users", Controllers.RegUser)
	app.Post("/users/login", Controllers.LoginUser)
	app.Delete("/users/:id", Controllers.DeleteUser)

	app.Listen("localhost:8080")
}
