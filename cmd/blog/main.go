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
		v1 = app.Group("/api/v1", api.JWTAuthmiddleware(mysqlConn))
	)

	//close db connection
	defer mysqlConn.Close()

	//v1 : post blog handler
	v1.Get("/posts", apiHandler.HandleGetPost)
	v1.Post("/posts", apiHandler.HandleInsertPost)
	v1.Get("/posts/:id", apiHandler.HandleGetPostById)
	v1.Get("/posts/title/:title", apiHandler.HandleGetPostByTitle)
	v1.Delete("/posts/:id", apiHandler.HandleDeletePost)
	v1.Put("/posts/:id", apiHandler.HandleUpdatePost)
	v1.Post("/posts/category", apiHandler.HandleInsertCategory)
	v1.Put("/posts/category/:id", apiHandler.HandleUpdateCategory)
	v1.Delete("/posts/category/:id", apiHandler.HandleDeleteCategory)
	v1.Get("/posts/category", apiHandler.HandleGetCategory)

	//admin router : user blog handler
	app.Get("/users", apiHandler.HandleGetUsers)
	app.Post("/users", apiHandler.HandleInsertUser)
	app.Delete("/users/:id", apiHandler.HandleDeleteUser)
	app.Put("/users/:id", apiHandler.HandleUpdateUser)

	app.Post("/users/signin", apiHandler.HandleSignInUser)
	app.Post("/users/signup", apiHandler.HandleSignUpUser)

	log.Fatal(app.Listen(":5000"))
}
