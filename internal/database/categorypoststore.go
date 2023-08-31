package database

import "log"

type CategoryStorer interface {
	InsertCategory([]int, int) error
	UpdateCategory(int, []int) error
	GetCountCategory(int) (int, error)
}

func (m *MysqlDatabase) InsertCategory(categoryID []int, postID int) error {
	query := `INSERT INTO CATEGORYPOST_TBL(categoryid, postid) VALUES (?,?);`
	stmt, err := m.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	for _, v := range categoryID {
		_, err = stmt.Exec(v, postID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *MysqlDatabase) DeleteCategory(postID int) error {
	query := `DELETE FROM CATEGORYPOST_TBL WHERE POSTID = ?`
	stmt, err := m.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(postID)
	if err != nil {
		return err
	}
	return nil
}

func (m *MysqlDatabase) UpdateCategory(postID int, categoryID []int) error {
	query := `UPDATE CATEGORYPOST_TBL SET CATEGORYID=? WHERE POSTID=?;`
	stmt, err := m.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	for _, v := range categoryID {
		log.Println(v, categoryID, postID)
		_, err = stmt.Exec(v, postID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *MysqlDatabase) GetCountCategory(postID int) (int, error) {
	query := `SELECT COUNT(*) FROM CATEGORYPOST_TBL WHERE POSTID = ?;`
	var count int
	err := m.DB.QueryRow(query, postID).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}
