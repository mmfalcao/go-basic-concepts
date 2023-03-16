package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3307)/cursogo")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// update
	stmt, _ := db.Prepare("update usuarios set nome = ? where id = ?")

	stmt.Exec("Julia", 4001)
	stmt.Exec("Larissa", 4000)

	// delete
	stmt2, _ := db.Prepare("delete from usuarios where id = ?")
	stmt2.Exec(4004)
}
