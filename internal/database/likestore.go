package database

import "fmt"

type LikeStorer interface {
	InsertLike(int, int) error
	DeleteLike(int, int) error
	GetLikesCount(int) (int, error)
	GetLikesCountByUserPostID(int, int) (int, error)
}

func (m *MysqlDatabase) InsertLike(userID, postID int) error {
	//check count of likes of postID for prevent duplicate like
	likesCount, err := m.GetLikesCountByUserPostID(userID, postID)
	if err != nil {
		return err
	}
	if likesCount == 1 {
		return fmt.Errorf("this post liked already ...")
	}

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
	res, err := stmt.Exec(userID, postID)
	if err != nil {
		return err
	}
	num, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if num == 0 {
		return fmt.Errorf("error to dislike this post! no rows affected ...")
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

func (m *MysqlDatabase) GetLikesCountByUserPostID(userID, postID int) (int, error) {
	query := `SELECT COUNT(*) FROM LIKES WHERE POSTID=? AND USERID=?;`
	var count int
	err := m.DB.QueryRow(query, postID, userID).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}
