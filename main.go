package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/moe-zdravstvo/moe-zdravstvo-main/database"
	"github.com/moe-zdravstvo/moe-zdravstvo-main/router"
)

func main() {
	database.InitMeilisearch()
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	router.InitRoutes(app)

	app.Static("/public", "./public/")
	app.Listen(":9990")
}
