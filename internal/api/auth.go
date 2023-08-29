package api

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/yousefzinsazk78/blog_json_api/internal/database"
)

func JWTAuthmiddleware(db database.UserStorer) fiber.Handler {
	return func(c *fiber.Ctx) error {
		jwtTokenString := c.Get("authToken")
		if jwtTokenString == "" {
			return NewBlogError(fiber.StatusBadRequest, "authToken")
		}
		res, err := ParseJWT(jwtTokenString)
		if err != nil {
			return NewBlogError(fiber.StatusBadRequest, "error in parsing jwt")
		}

		expiredAt, ok := res["ExpiredAt"].(float64)
		if !ok {
			return NewBlogError(fiber.StatusBadRequest, "is not ok to find expired at data")
		}
		intExpiredAt := int64(expiredAt)
		if time.Now().Unix() > intExpiredAt {
			return ErrTokenExpired()
		}
		userID, ok := res["userid"].(float64)
		if !ok {
			return NewBlogError(fiber.StatusBadRequest, "is not ok to find userid ")
		}
		intUserID := int(userID)
		user, err := db.GetUserByID(intUserID)
		if err != nil {
			return ErrUnAuthorized()
		}
		c.Context().SetUserValue("user", user)
		return c.Next()
	}

}
