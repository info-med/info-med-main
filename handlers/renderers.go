package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/moe-zdravstvo/moe-zdravstvo-main/database"
)

func RenderHomePage(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{}, "layouts/main")
}

func RenderSearchPage(c *fiber.Ctx) error {
	return c.Render("search", fiber.Map{})
}

func RenderGetDrugInfo(c *fiber.Ctx) error {
	id := c.Params("id")

	res := database.GetDrugInfo(id)

	return c.Render("drugInfo", res, "layouts/main")
}

func RenderAboutUsPage(c *fiber.Ctx) error {
	return c.Render("about", fiber.Map{}, "layouts/main")
}

func RenderUpgradePage(c *fiber.Ctx) error {
	return c.Render("upgrade", fiber.Map{}, "layouts/main")
}


func RenderAIPage(c *fiber.Ctx) error {
	return c.Render("ai", fiber.Map{}, "layouts/ai")
}
