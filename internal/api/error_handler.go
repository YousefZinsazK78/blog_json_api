package api

import (
	"github.com/gofiber/fiber/v2"
)

type ErrResp struct {
	Code int    `json:"status"`
	Msg  string `json:"message"`
}

func ErrorHandler(c *fiber.Ctx, err error) error {
	if error, ok := err.(BlogError); ok {
		return c.Status(error.Code).JSON(ErrResp{Code: error.Code, Msg: error.Message})
	}
	return c.Status(fiber.StatusBadRequest).JSON(ErrResp{Code: fiber.StatusBadRequest, Msg: "Bad Request"})
}

type BlogError struct {
	Code    int
	Message string
}

func (b BlogError) Error() string {
	return b.Message
}

func NewBlogError(code int, message string) BlogError {
	return BlogError{
		Code:    code,
		Message: message,
	}
}

func ErrPostBadRequest() BlogError {
	return BlogError{
		Code:    fiber.StatusBadRequest,
		Message: "Post Bad Request! 👎",
	}
}

func ErrNotFound() BlogError {
	return BlogError{
		Code:    fiber.StatusNotFound,
		Message: "Not Found! 🔎",
	}
}

func ErrTokenExpired() BlogError {
	return BlogError{
		Code:    fiber.StatusBadRequest,
		Message: "Token Expired! 😢",
	}
}

func ErrUnAuthorized() BlogError {
	return BlogError{
		Code:    fiber.StatusUnauthorized,
		Message: "unauthorized!",
	}
}

func ErrInternalServer() BlogError {
	return BlogError{
		Code:    fiber.StatusInternalServerError,
		Message: "internal server error!",
	}
}
