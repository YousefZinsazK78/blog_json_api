package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/yousefzinsazk78/blog_json_api/internal/types"
)

type PostStorer interface {
	InsertPost(*types.Post) error
	GetPosts(int, int) ([]*types.Post, error)
	GetPostsById(int) (*types.Post, error)
	GetPostByTitle(string) ([]*types.Post, error)
	DeletePost(int) error
	UpdatePost(int, types.UpdateParams) error
}

func (m *MysqlDatabase) InsertPost(postModel *types.Post) error {
	query := `INSERT INTO post_tbl(Title,Body, CreatedAt, UpdatedAt,user_id) VALUES (? ,? ,? ,?, ?)`
	stmt, err := m.DB.Prepare(query)
	if err != nil {
		return err
	}
	postModel.CreatedAt = time.Now().UTC()
	res, err := stmt.Exec(postModel.Title, postModel.Body, postModel.CreatedAt, postModel.UpdatedAt, postModel.AuthorID)
	if err != nil {
		return err
	}
	insertedID, err := res.LastInsertId()
	if err != nil {
		return err
	}

	if postModel.CategoryID != nil || len(postModel.CategoryID) != 0 {
		err = m.InsertCategory(postModel.CategoryID, int(insertedID))
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *MysqlDatabase) GetPosts(page, limit int) ([]*types.Post, error) {
	query := `SELECT * FROM POST_TBL LIMIT ?,?;`
	rows, err := m.DB.Query(query, (page-1)*limit, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []*types.Post
	for rows.Next() {
		var post types.Post
		if err := rows.Scan(&post.ID, &post.Title, &post.Body, &post.CreatedAt, &post.UpdatedAt, &post.AuthorID); err != nil {
			return nil, err
		}
		categoryModel, err := m.getCategory(post.ID)
		if err != nil {
			return nil, err
		}
		var cateID = []int{}
		for _, v := range categoryModel {
			cateID = append(cateID, v.ID)
		}
		post.CategoryID = cateID
		post.Likes, err = m.GetLikesCount(post.ID)
		if err != nil {
			return nil, err
		}
		post.Comments, err = m.GetCommentByPostID(post.ID)
		if err != nil {
			return nil, err
		}
		posts = append(posts, &post)
	}
	return posts, nil
}

func (m *MysqlDatabase) getCategory(postid int) ([]*types.Category, error) {
	query := `SELECT c.ID,c.NAME
	FROM CATEGORYPOST_TBL E LEFT JOIN CATEGORY_TBL c
	ON E.CATEGORYID= c.ID
	WHERE E.POSTID = ?;`
	rows, err := m.DB.Query(query, postid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var category []*types.Category
	for rows.Next() {
		var cat types.Category
		if err := rows.Scan(&cat.ID, &cat.CategoryName); err != nil {
			return nil, err
		}
		category = append(category, &cat)
	}
	return category, nil
}

func (m *MysqlDatabase) GetPostById(id int) (*types.Post, error) {
	const readvaluefromtbl = `SELECT * FROM post_tbl WHERE id=?;`
	res, err := m.DB.Query(readvaluefromtbl, id)
	if err != nil {
		return nil, err
	}
	defer res.Close()

	var post types.Post
	for res.Next() {
		err := res.Scan(&post.ID, &post.Title, &post.Body, &post.CreatedAt, &post.UpdatedAt, &post.AuthorID)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, fmt.Errorf("%d : post not found!", id)
			}
			return nil, err
		}
	}

	categoryModel, err := m.getCategory(post.ID)
	if err != nil {
		return nil, err
	}
	var cateID = []int{}
	for _, v := range categoryModel {
		cateID = append(cateID, v.ID)
	}
	post.CategoryID = cateID

	post.Likes, err = m.GetLikesCount(post.ID)
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (m *MysqlDatabase) GetPostByTitle(title string) ([]*types.Post, error) {
	query := "SELECT * FROM POST_TBL WHERE Title LIKE ?;"
	rows, err := m.DB.Query(query, "%"+title+"%")
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer rows.Close()

	var posts []*types.Post
	for rows.Next() {
		var post types.Post
		if err := rows.Scan(&post.ID, &post.Title, &post.Body, &post.CreatedAt, &post.UpdatedAt, &post.AuthorID, &post.CategoryID); err != nil {
			return nil, err
		}
		categoryModel, err := m.getCategory(post.ID)
		if err != nil {
			return nil, err
		}
		var cateID = []int{}
		for _, v := range categoryModel {
			cateID = append(cateID, v.ID)
		}
		post.CategoryID = cateID

		post.Likes, err = m.GetLikesCount(post.ID)
		if err != nil {
			return nil, err
		}
		posts = append(posts, &post)
	}

	return posts, nil
}

func (m *MysqlDatabase) DeletePost(id int) error {
	query := `DELETE FROM POST_TBL WHERE ID = ?`
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

func (m *MysqlDatabase) UpdatePost(id int, postParam *types.UpdateParams) error {
	query := `UPDATE POST_TBL SET title=?, updatedat=? WHERE ID = ?;`
	stmt, err := m.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(postParam.Title, time.Now().UTC(), id)
	if err != nil {
		return err
	}
	/// check there is category or not ....
	count, err := m.GetCountCategory(id)
	if err != nil {
		return err
	}

	if count > 0 {
		err := m.DeleteCategory(id)
		if err != nil {
			return err
		}
		err = m.InsertCategory(postParam.CategoryID, id)
		if err != nil {
			return err
		}
	} else {
		///insert category to categorypost_tbl
		err = m.InsertCategory(postParam.CategoryID, id)
		if err != nil {
			return err
		}
	}
	return nil
}
