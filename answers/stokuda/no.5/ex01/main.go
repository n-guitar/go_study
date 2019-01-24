package main

import (
	"database/sql"
	"fmt"
	"reflect"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	id         int
	username   string
	task       string
	end_flag   interface{}
	created_at string
	update_at  string
}

func main() {
	db, err := sql.Open("mysql", "go_user:go_user@/go_apps")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM tasks")
	if err != nil {
		panic(err.Error())
	}

	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}

	for _, column := range columns {
		fmt.Printf("%s |", column)
	}
	fmt.Println("")

	var user User
	for rows.Next() {
		err = rows.Scan(
			&user.id,
			&user.username,
			&user.task,
			&user.end_flag,
			&user.created_at,
			&user.update_at)

		if err != nil {
			panic(err.Error())
		}

		ref := reflect.ValueOf(user)
		for i := 0; i < ref.NumField(); i++ {
			fmt.Printf("%s |", ref.Field(i))
		}

		fmt.Println("")
	}

}
