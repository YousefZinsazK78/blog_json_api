package api

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/yousefzinsazk78/blog_json_api/internal/types"
)

func Admin(c *fiber.Ctx) error {
	user := c.Context().UserValue("user").(*types.User)
	// log.Println(user)
	if user == nil {
		log.Println("empty user")
		return ErrUnAuthorized()
	}
	if user.IsAdmin == false {
		log.Println("user is not admin")
		return ErrUnAuthorized()
	}
	return c.Next()
}
