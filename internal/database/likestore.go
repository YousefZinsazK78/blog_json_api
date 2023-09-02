package database

type LikeStorer interface {
	InsertLike(int, int) error
	DeleteLike(int, int) error
	GetLikesCount(int) (int, error)
}

func (m *MysqlDatabase) InsertLike(userID, postID int) error {
	query := `INSERT INTO LIKES(USERID, POSTID) VALUES(?, ?);`
	stmt, err := m.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(userID, postID)
	if err != nil {
		return err
	}
	return nil
}

func (m *MysqlDatabase) DeleteLike(userID, postID int) error {
	query := `DELETE FROM LIKES WHERE USERID = ? AND POSTID = ?;`
	stmt, err := m.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(userID, postID)
	if err != nil {
		return err
	}
	return nil
}

func (m *MysqlDatabase) GetLikesCount(postID int) (int, error) {
	query := `SELECT COUNT(*) FROM LIKES WHERE POSTID=?;`
	var count int
	err := m.DB.QueryRow(query, postID).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}
