package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, err := sql.Open(
		"sqlite3",
		"students.db",
	)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	rows, err := db.Query(
		"SELECT id, name, age FROM students ORDER BY id;",
	)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {

		var id int
		var name string
		var age int

		err := rows.Scan(
			&id,
			&name,
			&age,
		)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(
			id,
			name,
			age,
		)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
