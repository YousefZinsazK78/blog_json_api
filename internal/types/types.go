package types

import "time"

//Todo : implement comments section

type UpdateParams struct {
	Title string `json:"title"`
}

type Post struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
	// Author      User   `json:"author"`
	// Likes       []int  `json:"likes"`
	// Comments    []string  `json:"comments"`
	// Category  []string  `json:"category"`
	CreatedAt time.Time  `json:"createdat"`
	UpdatedAt *time.Time `json:"updatedat"`
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
