package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func RenderHomePage(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{})
}

func RenderSearchPage(c *fiber.Ctx) error {
	return c.Render("search", fiber.Map{})
}

func RenderAboutUsPage(c *fiber.Ctx) error {
	return c.Render("about", fiber.Map{})
}
