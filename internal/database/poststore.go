package database

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/yousefzinsazk78/blog_json_api/internal/types"
)

type PostStorer interface {
	InsertPost(*types.Post) error
	GetPosts() ([]*types.Post, error)
	GetPostsById(int) (*types.Post, error)
}

func (m *MysqlDatabase) InsertPost(postModel *types.Post) error {
	query := `INSERT INTO post_tbl(Title,Body, CreatedAt, UpdatedAt) VALUES (? ,? ,? ,?)`
	stmt, err := m.DB.Prepare(query)
	if err != nil {
		return err
	}
	postModel.CreatedAt = time.Now().UTC()
	_, err = stmt.Exec(postModel.Title, postModel.Body, postModel.CreatedAt, postModel.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (m *MysqlDatabase) GetPosts() ([]*types.Post, error) {
	query := `SELECT * FROM POST_TBL;`
	rows, err := m.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []*types.Post
	for rows.Next() {
		var post types.Post
		if err := rows.Scan(&post.ID, &post.Title, &post.Body, &post.CreatedAt, &post.UpdatedAt); err != nil {
			return nil, err
		}
		posts = append(posts, &post)
	}
	return posts, nil
}

func (m *MysqlDatabase) GetPostById(id int) (*types.Post, error) {
	// log.Println(id)
	// query := `SELECT * FROM post_tbl where ID = ?`
	// var post *types.Post
	// if err := m.DB.QueryRow(query, id).Scan(&post); err != nil {
	// 	log.Println("log in query row", post)
	// 	if err == sql.ErrNoRows {
	// 		return nil, fmt.Errorf("unknown id! : %d", id)
	// 	}
	// 	return nil, err
	// }
	// return post, nil
	const readvaluefromtbl = `SELECT * FROM post_tbl WHERE id=?;`
	res, err := m.DB.Query(readvaluefromtbl, id)
	if err != nil {
		return nil, err
	}
	defer res.Close()

	var post types.Post
	for res.Next() {
		err := res.Scan(&post.ID, &post.Title, &post.Body, &post.CreatedAt, &post.UpdatedAt)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, fmt.Errorf("%d : post not found!", id)
			}
			return nil, err
		}
	}
	return &post, nil
}
