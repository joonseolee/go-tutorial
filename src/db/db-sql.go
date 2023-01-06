package main

import (
	"database/sql"
	"fmt"
)

func main() {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, name FROM school")
	if err != nil {
		panic(err)
	}
	fmt.Println(rows)

	db.Exec("INSERT INTO school(id, name) VALUES (1, 'Alex')")
}
