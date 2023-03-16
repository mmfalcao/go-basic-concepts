package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
	return result
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3307)/")
	if err != nil {
		fmt.Println("#1")
		panic(err)
	}
	defer db.Close()

	for db.Ping() != nil {
		fmt.Println("Attempting connection to db")
		time.Sleep(5 * time.Second)
	}
	fmt.Println("Connected")

	exec(db, "create database if not exists cursogo")
	exec(db, "use cursogo")
	exec(db, "drop table if exists usuarios")
	exec(db, `create table usuarios (
		id integer auto_increment,
		nome varchar(80),
		PRIMARY KEY (id)
	)`)
}
