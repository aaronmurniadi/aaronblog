package controller

import (
	"github.com/aaronmurniadi/panacea/pages"
	"github.com/gofiber/fiber/v2"
)

func HomePageController(c *fiber.Ctx) error {
	return c.Type("html").SendString(pages.HomePage())
}
