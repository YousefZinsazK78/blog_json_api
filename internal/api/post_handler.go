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
		return ErrPostBadRequest()
	}

	err := a.mysqlDB.InsertPost(&post)
	if err != nil {
		return ErrPostBadRequest()
	}

	return c.Status(fiber.StatusOK).JSON(map[string]string{
		"message": "successfull inserted✅",
	})
}

func (a *Api) HandleGetPostById(c *fiber.Ctx) error {
	id := c.Params("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		return ErrPostBadRequest()
	}

	post, err := a.mysqlDB.GetPostById(intId)
	if err != nil {
		return ErrPostNotFound()
	}

	return c.Status(fiber.StatusOK).JSON(map[string]any{
		"result": post,
	})
}

func (a *Api) HandleGetPostByTitle(c *fiber.Ctx) error {
	title := c.Params("title")
	posts, err := a.mysqlDB.GetPostByTitle(title)
	if err != nil {
		return ErrPostNotFound()
	}

	return c.Status(fiber.StatusOK).JSON(map[string]any{
		"result": posts,
	})
}

func (a *Api) HandleDeletePost(c *fiber.Ctx) error {
	id := c.Params("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		return ErrPostBadRequest()
	}

	err = a.mysqlDB.DeletePost(intId)
	if err != nil {
		return ErrPostNotFound()
	}

	return c.Status(fiber.StatusOK).JSON(map[string]any{
		"status": fiber.StatusOK,
		"result": "post deleted successfully ✅",
	})
}

func (a *Api) HandleUpdatePost(c *fiber.Ctx) error {
	id := c.Params("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		return ErrPostBadRequest()
	}

	var post *types.UpdateParams
	if err := c.BodyParser(post); err != nil {
		return ErrPostBadRequest()
	}

	err = a.mysqlDB.UpdatePost(intId, post)
	if err != nil {
		return ErrPostNotFound()
	}

	return c.Status(fiber.StatusOK).JSON(map[string]any{
		"status": fiber.StatusOK,
		"result": "post deleted successfully ✅",
	})
}
