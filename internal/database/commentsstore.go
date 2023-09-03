package database

import (
	"fmt"
	"log"
	"time"

	"github.com/yousefzinsazk78/blog_json_api/internal/types"
)

type CommentsStorer interface {
	InsertComment(*types.Comment) error
	DeleteComment(*types.DeleteComments) error
}

func (m *MysqlDatabase) InsertComment(comment *types.Comment) error {
	query := `INSERT INTO COMMENTS_TBL(content, postid, userid, createdAt) VALUES (?,?,?,?);`
	stmt, err := m.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.Exec(comment.Content, comment.PostID, comment.UserID, time.Now().UTC())
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("error : rows affected : %s", err.Error())
	}
	return nil
}

func (m *MysqlDatabase) DeleteComment(comment *types.DeleteComments) error {
	log.Println(comment)
	query := `DELETE FROM COMMENTS_TBL WHERE USERID=? AND POSTID=?;`
	stmt, err := m.DB.Prepare(query)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	defer stmt.Close()
	res, err := stmt.Exec(comment.UserID, comment.PostID)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	if rowsAffected == 0 {
		log.Println(err.Error())
		return fmt.Errorf("error : rows affected : %s", err.Error())
	}
	return nil
}
