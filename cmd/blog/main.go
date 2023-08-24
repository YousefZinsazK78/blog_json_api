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
		app        = fiber.New()
		apiHandler = api.New()
		mysqlConn  = database.NewMysqlConn(os.Getenv("Username"), os.Getenv("Password"), os.Getenv("Net"), os.Getenv("Addr"), os.Getenv("DBName"))
	)

	//close db connection
	defer mysqlConn.Close()

	//get method all posts
	app.Get("/", apiHandler.HandleIndex)

	log.Fatal(app.Listen(":5000"))
}
