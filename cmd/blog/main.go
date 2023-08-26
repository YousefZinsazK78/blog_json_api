package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/yousefzinsazk78/blog_json_api/internal/api"
	"github.com/yousefzinsazk78/blog_json_api/internal/database"
)

func main() {
	err := godotenv.Load("./internal/configs/.env")
	if err != nil {
		log.Fatal("Error: unable to load .env file")
	}

	var (
		mysqlConn = database.NewMysqlConn(os.Getenv("Username"), os.Getenv("Password"), os.Getenv("Net"), os.Getenv("Addr"), os.Getenv("DBName"))

		apiHandler = api.New(*mysqlConn)
		app        = fiber.New(
			fiber.Config{
				ErrorHandler: api.ErrorHandler,
			},
		)
		v1 = app.Group("/api/v1")
	)

	//close db connection
	defer mysqlConn.Close()

	//v1 : post blog handler
	v1.Get("/posts", apiHandler.HandleGetPost)
	v1.Post("/posts", apiHandler.HandleInsertPost)
	v1.Get("/posts/:id", apiHandler.HandleGetPostById)
	v1.Get("/posts/title/:title", apiHandler.HandleGetPostByTitle)

	log.Fatal(app.Listen(":5000"))
}
