package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/moe-zdravstvo/moe-zdravstvo-main/handlers"
)

func InitRoutes(app *fiber.App) {
	// Main Routes
	app.Get("/search", handlers.HandleSearch)
	//app.Get("/drugInfo", handlers.HandleGetDrugInfo)
	app.Get("/drugstoreInfo", handlers.HandleGetDrugstoreInfo)

	// Rendering Routes
	app.Get("/", handlers.RenderHomePage)
	app.Get("/get/drug/:id", handlers.RenderGetDrugInfo)
	app.Get("/about", handlers.RenderAboutUsPage)
	app.Get("/upgrade", handlers.RenderUpgradePage)
	app.Get("/sources", handlers.RenderSourcesPage)
	app.Get("/get/drugstore/:id", handlers.RenderDrugstorePage)

	// AI
	app.Get("/ai", handlers.RenderAIPage)
}
