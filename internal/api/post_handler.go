package api

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/yousefzinsazk78/blog_json_api/internal/types"
)

func (a *Api) HandleGetPost(c *fiber.Ctx) error {
	posts, err := a.mysqlDB.GetPosts()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("error to load post list")
	}
	return c.Status(fiber.StatusOK).JSON(map[string]any{
		"size":    len(posts),
		"message": posts,
	})
}

func (a *Api) HandleInsertPost(c *fiber.Ctx) error {
	var post types.Post
	if err := c.BodyParser(&post); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	err := a.mysqlDB.InsertPost(&post)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("error to load post list")
	}

	return c.Status(fiber.StatusOK).JSON(map[string]string{
		"message": "successfull inserted✅",
	})
}

func (a *Api) HandleGetPostById(c *fiber.Ctx) error {
	id := c.Params("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]string{
			"message": "unable to convert asci to integer",
		})
	}

	post, err := a.mysqlDB.GetPostById(intId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("error to load post ")
	}

	return c.Status(fiber.StatusOK).JSON(map[string]any{
		"result": post,
	})
}
