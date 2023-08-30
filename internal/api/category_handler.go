package api

// import (
// 	"log"
// 	"strconv"

// 	"github.com/gofiber/fiber/v2"
// 	"github.com/yousefzinsazk78/blog_json_api/internal/types"
// )

// func (a *Api) HandleInsertCategory(c *fiber.Ctx) error {
// 	var category types.Category
// 	if err := c.BodyParser(&category); err != nil {
// 		return NewBlogError(fiber.StatusBadRequest, err.Error())
// 	}

// 	err := a.mysqlDB.InsertCategory(&category)
// 	if err != nil {
// 		return NewBlogError(fiber.StatusBadRequest, err.Error())
// 	}

// 	return c.Status(fiber.StatusCreated).JSON(types.Response{
// 		Status:  fiber.StatusCreated,
// 		Message: "successfully inserted",
// 	})
// }

// func (a *Api) HandleUpdateCategory(c *fiber.Ctx) error {
// 	id := c.Params("id")
// 	intID, err := strconv.Atoi(id)
// 	if err != nil {
// 		return ErrInternalServer()
// 	}
// 	var category types.Category
// 	if err := c.BodyParser(&category); err != nil {
// 		return NewBlogError(fiber.StatusBadRequest, err.Error())
// 		// return ErrPostBadRequest()
// 	}

// 	err = a.mysqlDB.UpdateCategory(intID, &category)
// 	if err != nil {
// 		return ErrInternalServer()
// 	}

// 	return c.Status(fiber.StatusOK).JSON(types.Response{
// 		Status:  fiber.StatusOK,
// 		Message: "successfully updated...",
// 	})
// }

// func (a *Api) HandleDeleteCategory(c *fiber.Ctx) error {
// 	id := c.Params("id")
// 	intID, err := strconv.Atoi(id)
// 	if err != nil {
// 		return ErrInternalServer()
// 	}

// 	err = a.mysqlDB.DeleteCategory(intID)
// 	if err != nil {
// 		return ErrInternalServer()
// 	}

// 	return c.Status(fiber.StatusOK).JSON(types.Response{
// 		Status:  fiber.StatusOK,
// 		Message: "successfully deleted...",
// 	})
// }

// func (a *Api) HandleGetCategory(c *fiber.Ctx) error {
// 	var queryParams types.QueryParams
// 	if err := c.QueryParser(&queryParams); err != nil {
// 		return NewBlogError(fiber.StatusBadRequest, err.Error())
// 	}
// 	log.Println(queryParams)
// 	categories, err := a.mysqlDB.GetCategory(queryParams.Pages, queryParams.Limits)
// 	if err != nil {
// 		log.Println(err)
// 		return NewBlogError(fiber.StatusBadRequest, err.Error())
// 	}

// 	return c.Status(fiber.StatusOK).JSON(types.Response{
// 		Status:  fiber.StatusOK,
// 		Message: categories,
// 	})
// }
