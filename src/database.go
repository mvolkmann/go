package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql" // anonymous import
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	db, err := sql.Open("mysql",
		"root:root@tcp(127.0.0.1:3306)/tour_of_heroes")
	check(err)
	defer db.Close()

	rows, err := db.Query("select name from hero")
	check(err)
	defer rows.Close()

	var name string
	for rows.Next() {
		err := rows.Scan(&name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("name =", name)
	}
	err = rows.Err()
	check(err)
}
