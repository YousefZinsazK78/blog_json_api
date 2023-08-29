package database

import "github.com/yousefzinsazk78/blog_json_api/internal/types"

type CategoryStorer interface {
	InsertCategory(*types.Category) error
	UpdateCategory(int, *types.Category) error
	DeleteCategory(int) error
	GetCategory(int, int) ([]*types.Category, error)
}

func (m *MysqlDatabase) InsertCategory(category *types.Category) error {
	query := `INSERT INTO CATEGORY_TBL(name) VALUES (?);`
	stmt, err := m.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(category.CategoryName)
	if err != nil {
		return err
	}
	return nil
}

func (m *MysqlDatabase) UpdateCategory(id int, category *types.Category) error {
	query := `UPDATE CATEGORY_TBL SET name=? WHERE ID=?;`
	stmt, err := m.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(category.CategoryName, id)
	if err != nil {
		return err
	}
	return nil
}

func (m *MysqlDatabase) DeleteCategory(id int) error {
	query := `DELETE FROM CATEGORY_TBL WHERE ID=?;`
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

func (m *MysqlDatabase) GetCategory(pages, limits int) ([]*types.Category, error) {
	query := `SELECT * FROM CATEGORY_TBL LIMIT ?,?;`
	rows, err := m.DB.Query(query, (pages-1)*limits, limits)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []*types.Category
	for rows.Next() {
		var category types.Category
		if err := rows.Scan(&category.ID, &category.CategoryName); err != nil {
			return nil, err
		}
		categories = append(categories, &category)
	}
	return categories, nil
}
