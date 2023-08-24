package database

import "github.com/yousefzinsazk78/blog_json_api/internal/types"

type PostStorer interface {
	InsertPost(*types.Post) error
}

func (m *MysqlDatabase) InsertPost(postModel *types.Post) error {
	query := `INSERT INTO postTbl(Title,Description, Author, Likes, CreatedAt, UpdatedAt) VALUES (?, ? ,? ,? ,? ,?)`
	stmt, err := m.DB.Prepare(query)
	if err != nil {
		return err
	}
	_ = stmt
	return nil
}
