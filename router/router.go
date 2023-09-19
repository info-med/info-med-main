package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/moe-zdravstvo/moe-zdravstvo-main/handlers"
)

func InitRoutes(app *fiber.App) {
	// (TODO: this page routing/rendering should somehow be done with layouts not like this)

	// Main Routes
	app.Get("/search", handlers.HandleSearch)
	app.Get("/drugInfo", handlers.HandleGetDrugInfo)
	app.Get("/drugstoreInfo", handlers.HandleGetDrugstoreInfo)

	// Rendering Routes
	app.Get("/", handlers.RenderHomePage)
	app.Get("/renderSearch", handlers.RenderSearchPage)
	app.Get("/renderAboutUs", handlers.RenderAboutUsPage)
}
