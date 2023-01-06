package main

import (
	"database/sql"
	"fmt"
	"log"
)

func main() {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var name string
	err = db.QueryRow("SELECT name FROM test1 WHERE id = 1").Scan(&name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(name)

	var id int
	var schoolName string
	rows, err := db.Query("SELECT id, name FROM school WHERE id >= ?", 1)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &schoolName)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println()
	}
}
