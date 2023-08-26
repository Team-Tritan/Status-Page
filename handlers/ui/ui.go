package handlers

import "github.com/gofiber/fiber/v2"

func DisplayUI(c *fiber.Ctx) error {
	return c.Render("./pages/index.html", fiber.Map{})
}
