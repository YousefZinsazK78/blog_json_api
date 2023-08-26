package api

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/yousefzinsazk78/blog_json_api/internal/types"
)

func (a *Api) HandleInsertUser(c *fiber.Ctx) error {
	var user types.User
	if err := c.BodyParser(&user); err != nil {
		return ErrPostBadRequest()
	}

	log.Println(user.HashUserPassword())
	user.Password = user.HashUserPassword()

	err := a.mysqlDB.InsertUser(&user)
	if err != nil {
		return ErrPostBadRequest()
	}
	return c.Status(fiber.StatusCreated).JSON(map[string]any{
		"message": "user inserted successfully ✅",
	})
}
