package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "go_user:go_user@tcp(localhost:3306)/go_apps")
	if err != nil {
		fmt.Println("DB Connection Error")
	}
	defer db.Close()

	rows, _ := db.Query("SELECT * FROM tasks")
	col, _ := rows.Columns()

	colindex := make([]sql.RawBytes, len(col))
	colvals := make([]interface{}, len(colindex))

	for i := range colindex {
		colvals[i] = &colindex[i]
	}
	sqlhandler := func(w http.ResponseWriter, req *http.Request) {
		var val string

		for rows.Next() {

			for i, col := range colindex {
				val = string(col)
				io.WriteString(w, col[i])
				io.WriteString(w, ":")
				io.WriteString(w, val)
				io.WriteString(w, "\n")
			}
			io.WriteString(w, "--------------------------------------\n")
		}
	}
	http.HandleFunc("/pretty", sqlhandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}