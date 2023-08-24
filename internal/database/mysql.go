package database

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

type MysqlDatabase struct {
	*sql.DB
}

func NewMysqlConn(username, password, net, addr, dbname string) *MysqlDatabase {
	config := mysql.Config{
		User:   username,
		Passwd: password,
		Net:    net,
		Addr:   addr,
		DBName: dbname,
	}
	//create database connection
	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		log.Fatal("error: unable to open mysql database connection")
		return nil
	}

	return &MysqlDatabase{
		DB: db,
	}
}
