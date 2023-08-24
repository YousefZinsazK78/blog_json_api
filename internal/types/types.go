package types

import "time"

//Todo : implement comments section

type Post struct {
	ID          int    `json:"-"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      Author `json:"author"`
	Likes       []int  `json:"likes"`
	// Comments    []string  `json:"comments"`
	CreatedAt time.Time `json:"createdat"`
	UpdatedAt time.Time `json:"updatedat"`
}

type Author struct {
	Fullname string
	Username string
	Password string
}

type User struct {
	Fullname string
	Username string
	Password string
	email    string
	role     Roles
}

type Roles int

const (
	Admin  Roles = 1
	Writer Roles = 2
	Guest  Roles = 3
)
