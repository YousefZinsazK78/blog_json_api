package api

import (
	"log"
	"strconv"

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
	return c.Status(fiber.StatusCreated).JSON(
		types.Response{
			Status:  fiber.StatusCreated,
			Message: "inserted successfully ✅",
		},
	)
}

func (a *Api) HandleDeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		return ErrPostBadRequest()
	}

	err = a.mysqlDB.DeleteUser(intId)
	if err != nil {
		return ErrPostBadRequest()
	}
	return c.Status(fiber.StatusAccepted).JSON(
		types.Response{
			Status:  fiber.StatusAccepted,
			Message: "deleted successfully ✅",
		},
	)
}

func (a *Api) HandleUpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		return ErrPostBadRequest()
	}
	var userUpdateParams types.UserUpdateParams
	if err := c.BodyParser(&userUpdateParams); err != nil {
		return ErrPostBadRequest()
	}
	//todo : update user
	err = a.mysqlDB.UpdateUser(intId, &userUpdateParams)
	if err != nil {
		return ErrPostBadRequest()
	}
	return c.Status(fiber.StatusAccepted).JSON(
		types.Response{
			Status:  fiber.StatusAccepted,
			Message: "deleted successfully ✅",
		},
	)
}

func (a *Api) HandleGetUsers(c *fiber.Ctx) error {
	var queryParams types.QueryParams
	if err := c.QueryParser(&queryParams); err != nil {
		return ErrPostBadRequest()
	}

	users, err := a.mysqlDB.GetUsers(queryParams.Pages, queryParams.Limits)
	if err != nil {
		return ErrPostBadRequest()
	}

	return c.Status(fiber.StatusAccepted).JSON(
		types.Response{
			Status:  fiber.StatusAccepted,
			Message: users,
		},
	)
}

func (a *Api) HandleSignInUser(c *fiber.Ctx) error {
	var SignUserParams types.UserSignInParams
	if err := c.BodyParser(&SignUserParams); err != nil {
		return ErrPostBadRequest()
	}

	user, err := a.mysqlDB.GetUserByEmail(SignUserParams.Email)
	if err != nil {
		return ErrNotFound()
	}

	var resp types.Response

	if user.CheckHashPassword(SignUserParams.Password) {
		resp = types.Response{
			Status:  fiber.StatusAccepted,
			Message: "signIn successfully 🤘✅",
		}
	} else {
		resp = types.Response{
			Status:  fiber.StatusBadRequest,
			Message: "signIn problem ⚠",
		}
	}

	return c.Status(resp.Status).JSON(
		resp,
	)
}
