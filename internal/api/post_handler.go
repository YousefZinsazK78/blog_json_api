package api

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/yousefzinsazk78/blog_json_api/internal/types"
)

func (a *Api) HandleGetPost(c *fiber.Ctx) error {
	var queryParams types.QueryParams
	if err := c.QueryParser(&queryParams); err != nil {
		log.Println(err)
		return ErrPostBadRequest()
	}

	if queryParams.Pages == 0 {
		queryParams.Pages = 1
	}

	posts, err := a.mysqlDB.GetPosts(queryParams.Pages, queryParams.Limits)
	if err != nil {
		return NewBlogError(fiber.StatusBadRequest, err.Error())
		// return ErrPostBadRequest()
	}
	return c.Status(fiber.StatusOK).JSON(map[string]any{
		"length":  len(posts),
		"message": posts,
	})
}

func (a *Api) HandleInsertPost(c *fiber.Ctx) error {
	var post types.Post
	if err := c.BodyParser(&post); err != nil {
		return NewBlogError(fiber.StatusBadRequest, err.Error())
	}

	err := a.mysqlDB.InsertPost(&post)
	if err != nil {
		return NewBlogError(fiber.StatusBadRequest, err.Error())
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
		return ErrNotFound()
	}

	return c.Status(fiber.StatusOK).JSON(map[string]any{
		"result": post,
	})
}

func (a *Api) HandleGetPostByTitle(c *fiber.Ctx) error {
	title := c.Params("title")
	posts, err := a.mysqlDB.GetPostByTitle(title)
	if err != nil {
		return ErrNotFound()
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
		log.Println(err.Error())
		return ErrInternalServer()
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

	var post types.UpdateParams
	if err := c.BodyParser(&post); err != nil {
		return NewBlogError(fiber.StatusBadRequest, err.Error())
	}

	err = a.mysqlDB.UpdatePost(intId, &post)
	if err != nil {
		return ErrNotFound()
	}

	return c.Status(fiber.StatusOK).JSON(map[string]any{
		"status": fiber.StatusOK,
		"result": "post updated successfully ✅",
	})
}

func (a *Api) HandleLikesPost(c *fiber.Ctx) error {
	var likePost types.LikesParams
	if err := c.BodyParser(&likePost); err != nil {
		return NewBlogError(fiber.StatusBadRequest, err.Error())
	}

	err := a.mysqlDB.InsertLike(likePost.UserID, likePost.PostID)
	if err != nil {
		return ErrNotFound()
	}

	return c.Status(fiber.StatusCreated).JSON(map[string]any{
		"status": fiber.StatusCreated,
		"result": "post liked 👍",
	})
}

func (a *Api) HandleDisLikesPost(c *fiber.Ctx) error {
	var likePost types.LikesParams
	if err := c.BodyParser(&likePost); err != nil {
		return NewBlogError(fiber.StatusBadRequest, err.Error())
	}

	err := a.mysqlDB.DeleteLike(likePost.UserID, likePost.PostID)
	if err != nil {
		return ErrNotFound()
	}

	return c.Status(fiber.StatusCreated).JSON(map[string]any{
		"status": fiber.StatusCreated,
		"result": "post disliked 👎",
	})
}
