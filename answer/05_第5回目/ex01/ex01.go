package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// わからなすぎて、鼻血でそうｗｗｗｗ　2019/12/3　22時37分のことである。

func main() {

	ryo, err := sql.Open("mysql", "go_user:go_user@tcp(localhost:3306)/go_apps")
	if err != nil {
		panic(err.Error())
	}
	defer ryo.Close()

	rows, err := ryo.Query("SELECT * FROM tasks")
	if err != nil {
		panic(err.Error())
	}

	kon, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}

	values := make([]sql.RawBytes, len(kon))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}

		var value string
		for i, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			fmt.Println(kon[i], ": ", value)
		}
		fmt.Println("-----------------------------------")
	}
}

/*
	for rows.Next() {
		var id int
		var name string

		if err := rows.Scan(&id, &name); err != nil {
			panic(err.Error())
		}
		fmt.Println(id, name)
	}
*/
