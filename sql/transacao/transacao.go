package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3307)/cursogo")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("insert into usuarios(id, nome) values(?,?)")

	stmt.Exec(4000, "Bia")
	stmt.Exec(4001, "Carlos")
	_, err = stmt.Exec(1, "Tiago") // chave duplicada

	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	tx.Commit()
}
