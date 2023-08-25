package api

import "github.com/yousefzinsazk78/blog_json_api/internal/database"

type Api struct {
	mysqlDB database.MysqlDatabase
}

func New(mysqlDB database.MysqlDatabase) *Api {
	return &Api{
		mysqlDB: mysqlDB,
	}
}
