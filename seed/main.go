package main

import (
	"database/sql"
	"log"
	"time"

	"github.com/joho/godotenv"
	"github.com/yousefzinsazk78/blog_json_api/internal/types"
)

func seedCategoryTable(db *sql.DB, categoryname string, postid int) {
	query := `INSERT INTO category_tbl(NAME, postid) VALUES (? ,?)`
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(categoryname, postid)
	if err != nil {
		log.Fatal(err)
	}
}

func seedCategoryPostTable(db *sql.DB, categoryid, postid int) {
	query := `INSERT INTO categorypost_tbl(categoryid, postid) VALUES (? ,?)`
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(categoryid, postid)
	if err != nil {
		log.Fatal(err)
	}
}

func seedPostTable(db *sql.DB, title, body string, userid int) {
	query := `INSERT INTO post_tbl(Title,Body, CreatedAt,user_id) VALUES (? ,? ,?, ?)`
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(title, body, time.Now().UTC(), userid)
	if err != nil {
		log.Fatal(err)
	}
}

func seedUserTable(db *sql.DB, fullname, email, username, password string, isAdmin bool) {
	query := `INSERT INTO user_tbl(fullname, email, username, isAdmin , password) VALUES (?,?,?,?,?);`
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	var user = types.User{}
	user.Password = password
	hashPasswd := user.HashUserPassword()
	_, err = stmt.Exec(fullname, email, username, isAdmin, hashPasswd)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	err := godotenv.Load("./internal/configs/.env")
	if err != nil {
		log.Fatal("Error: unable to load .env file")
	}

	// mysqlConn := database.NewMysqlConn(os.Getenv("Username"), os.Getenv("Password"), os.Getenv("Net"), os.Getenv("Addr"), os.Getenv("DBName"))
	// seedCategoryTable(mysqlConn.DB, "science", 2505)
	// seedCategoryTable(mysqlConn.DB, "science", 2504)
	// seedCategoryTable(mysqlConn.DB, "sport", 2505)
	// seedCategoryTable(mysqlConn.DB, "sport", 2504)
	// var wg sync.WaitGroup

	// wg.Add(1)
	// go func() {
	// 	for i := 0; i < 15; i++ {

	// 		// seedPostTable(mysqlConn.DB, "title test blog post haha", "body test blog post", 20)
	// 	}
	// 	wg.Done()
	// }()

	// wg.Wait()

	// go func() {
	// 	for i := 0; i < 250; i++ {
	// 		seedPostTable(mysqlConn.DB, "blog post title", "blog post body", 19)
	// 	}
	// 	wg.Done()
	// }()
	// go func() {
	// 	for i := 0; i < 250; i++ {
	// 		seedPostTable(mysqlConn.DB, "test test title test test", "test body test", 19)
	// 	}
	// 	wg.Done()
	// }()

	// wg.Wait()

	// seedUserTable(mysqlConn.DB, "yousef kashani", "mina@email.com", "yousefusername", "password123", true)
	// seedUserTable(mysqlConn.DB, "mina irani", "mina@email.com", "minausername", "password123", false)

}
