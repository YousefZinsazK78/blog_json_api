package main

import (
	"database/sql"
	"log"
	"os"
	"sync"
	"time"

	"github.com/joho/godotenv"
	"github.com/yousefzinsazk78/blog_json_api/internal/database"
)

func seedPostTable(db *sql.DB, title, body string) {
	query := `INSERT INTO post_tbl(Title,Body, CreatedAt) VALUES (? ,? ,?)`
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(title, body, time.Now().UTC())
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	err := godotenv.Load("./internal/configs/.env")
	if err != nil {
		log.Fatal("Error: unable to load .env file")
	}

	mysqlConn := database.NewMysqlConn(os.Getenv("Username"), os.Getenv("Password"), os.Getenv("Net"), os.Getenv("Addr"), os.Getenv("DBName"))
	var wg sync.WaitGroup

	wg.Add(3)
	go func() {
		for i := 0; i < 250; i++ {
			seedPostTable(mysqlConn.DB, "title test blog post", "body test blog post")
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 250; i++ {
			seedPostTable(mysqlConn.DB, "blog post title", "blog post body")
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 250; i++ {
			seedPostTable(mysqlConn.DB, "test test title test test", "test body test")
		}
		wg.Done()
	}()

	wg.Wait()
}
