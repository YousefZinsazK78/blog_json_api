package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yousefzinsazk78/blog_json_api/internal/types"
)

func (a *Api) HandleCommentsPost(c *fiber.Ctx) error {
	var comment types.Comment
	if err := c.BodyParser(&comment); err != nil {
		return ErrPostBadRequest()
	}

	user := c.Context().UserValue("user").(*types.User)
	comment.UserID = user.ID

	err := a.mysqlDB.InsertComment(&comment)
	if err != nil {
		return ErrInternalServer()
	}
	return c.Status(fiber.StatusAccepted).JSON(
		types.Response{
			Status:  fiber.StatusAccepted,
			Message: "comment added 📝",
		},
	)
}

func (a *Api) HandleCommentsDelete(c *fiber.Ctx) error {
	var comment types.DeleteComments
	if err := c.BodyParser(&comment); err != nil {
		return ErrPostBadRequest()
	}

	user := c.Context().UserValue("user").(*types.User)
	comment.UserID = user.ID

	err := a.mysqlDB.DeleteComment(&comment)
	if err != nil {
		return ErrInternalServer()
	}
	return c.Status(fiber.StatusAccepted).JSON(
		types.Response{
			Status:  fiber.StatusAccepted,
			Message: "comment deleted 📝",
		},
	)
}
