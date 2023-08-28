package database

import (
	"github.com/yousefzinsazk78/blog_json_api/internal/types"
)

type UserStorer interface {
	GetUsers(int, int) ([]*types.User, error)
	GetUserByEmail(string) (*types.User, error)
	GetUserByID(int) (*types.User, error)
	InsertUser(*types.User) error
	DeleteUser(int) error
	UpdateUser(int, *types.UserUpdateParams) error
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

func (m *MysqlDatabase) DeleteUser(id int) error {
	query := `DELETE FROM USER_TBL WHERE ID = ?;`
	stmt, err := m.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

func (m *MysqlDatabase) UpdateUser(id int, user *types.UserUpdateParams) error {
	query := `UPDATE USER_TBL SET username=?, email=? WHERE ID=?;`
	stmt, err := m.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.Username, user.Email, id)
	if err != nil {
		return err
	}
	return nil
}

func (m *MysqlDatabase) GetUsers(page, limit int) ([]*types.User, error) {
	query := `SELECT * FROM USER_TBL LIMIT ?,?;`
	rows, err := m.DB.Query(query, (page-1)*limit, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*types.User
	for rows.Next() {
		var user types.User
		if err := rows.Scan(&user.ID, &user.Fullname, &user.Email, &user.Username, &user.IsAdmin, &user.Password); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

func (m *MysqlDatabase) GetUserByEmail(email string) (*types.User, error) {
	query := `SELECT * FROM USER_TBL WHERE EMAIL=?;`
	var user types.User
	if err := m.DB.QueryRow(query, email).Scan(&user.ID, &user.Fullname, &user.Email, &user.Username, &user.IsAdmin, &user.Password); err != nil {
		return nil, err
	}
	return &user, nil
}

func (m *MysqlDatabase) GetUserByID(id int) (*types.User, error) {
	query := `SELECT * FROM USER_TBL WHERE ID=?;`
	var user types.User
	if err := m.DB.QueryRow(query, id).Scan(&user.ID, &user.Fullname, &user.Email, &user.Username, &user.IsAdmin, &user.Password); err != nil {
		return nil, err
	}
	return &user, nil
}
