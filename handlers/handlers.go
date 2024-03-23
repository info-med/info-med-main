package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/moe-zdravstvo/moe-zdravstvo-main/database"
)

func HandleSearch(c *fiber.Ctx) error {
	query := c.Query("search")

	if query != "" {
		res := database.Search(query)

		return c.Render("searchResult", res)
	}

	return c.Status(fiber.StatusBadRequest).JSON(map[string]string{"error": "No search query provided"})
}

func HandleGetDrugstoreInfo(c *fiber.Ctx) error {
	query := c.Query("drugstoreId")

	if query != "" {
		res := database.GetDrugstoreInfo(query)

		return c.Render("drugstoreModal", res)
	}

	return c.Status(fiber.StatusBadRequest).JSON(map[string]string{"error": "No drugstore ID provided"})
}
