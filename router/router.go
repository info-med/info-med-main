package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/moe-zdravstvo/moe-zdravstvo-main/handlers"
)

func InitRoutes(app *fiber.App) {
	// Main Routes
	app.Get("/search", handlers.HandleSearch)
	app.Get("/drugInfo", handlers.HandleGetDrugInfo)
	app.Get("/drugstoreInfo", handlers.HandleGetDrugstoreInfo)

	// Rendering Routes
	app.Get("/", handlers.RenderHomePage)
	app.Get("/renderSearch", handlers.RenderSearchPage)
	app.Get("/about", handlers.RenderAboutUsPage)

  // AI
  app.Get("/ai", handlers.RenderAIPage)
}
