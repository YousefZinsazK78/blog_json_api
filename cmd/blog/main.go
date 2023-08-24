package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/yousefzinsazk78/blog_json_api/internal/api"
)

func main() {
	app := fiber.New()
	apiHandler := api.New()

	//get method all posts
	app.Get("/", apiHandler.HandleIndex)

	log.Fatal(app.Listen(":5000"))
}
