package api

import "github.com/gofiber/fiber/v2"

func (a *Api) HandleIndex(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).SendString("this is get method and this is response")
}
