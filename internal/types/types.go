package types

import (
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type LikesParams struct {
	PostID int `json:"postid"`
	UserID int `json:"userid"`
}

type UserSignInParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Response struct {
	Status  int `json:"status"`
	Message any `json:"message"`
}

type UserUpdateParams struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type QueryParams struct {
	Pages  int `json:"pages"`
	Limits int `json:"limits"`
}

type UpdateParams struct {
	Title      string `json:"title"`
	CategoryID []int  `json:"categoryid"`
}

type Category struct {
	ID           int    `json:"id"`
	CategoryName string `json:"categoryname"`
}

type Post struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Body     string `json:"body"`
	AuthorID int    `json:"authorid"`
	// Likes       []int  `json:"likes"`
	// Comments    []string  `json:"comments"`
	CategoryID []int      `json:"categoryid"`
	CreatedAt  time.Time  `json:"createdat"`
	UpdatedAt  *time.Time `json:"updatedat"`
}

type User struct {
	ID       int    `json:"id"`
	Fullname string `json:"fullname"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	IsAdmin  bool   `json:"-"`
}

func (u User) HashUserPassword() string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	if err != nil {
		log.Fatal(err)
	}
	return string(bytes)
}

func (u *User) CheckHashPassword(Passwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(Passwd))
	return err == nil
}
