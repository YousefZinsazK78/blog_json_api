package database

type CategoryStorer interface {
	InsertCategory([]int, int) error
}

func (m *MysqlDatabase) InsertCategoryPost(categoryID []int, postID int) error {
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
