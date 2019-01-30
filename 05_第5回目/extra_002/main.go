package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func selectSQL(db *sql.DB) {
	rows, err := db.Query("select * from tasks")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id string
		var username string
		err := rows.Scan(&id, &username)
		if err != nil {
			panic(err)
		}
		fmt.Printf("id:%s\t username:%s\n", id, username)
	}
}

func main() {
	db, err := sql.Open("mysql", "go_user:go_user@tcp(localhost:3306)/go_apps")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	selectSQL(db)
}
