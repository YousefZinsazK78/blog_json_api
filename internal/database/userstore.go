package database

import (
	"github.com/yousefzinsazk78/blog_json_api/internal/types"
)

type UserStorer interface {
	InsertUser(*types.User) error
}

func (m *MysqlDatabase) InsertUser(user *types.User) error {
	query := `INSERT INTO user_tbl(fullname, email, username, password) VALUES (?,?,?,?);`
	stmt, err := m.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Fullname, user.Email, user.Username, user.Password)
	if err != nil {
		return err
	}
	return nil
}
