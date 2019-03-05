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

	rows, err := db.Query("SELECT * FROM tasks")
	if err != nil {
		fmt.Println("Unable to select from table")
	}

	clm, err := rows.Columns()
	if err != nil {
		fmt.Println("Unable to fetch colums from table")
	}

	clmbox := make([]sql.RawBytes, len(clm))
	clmvals := make([]interface{}, len(clmbox))

	for i := range clmbox {
		clmvals[i] = &clmbox[i]
	}
	sqlhandler := func(w http.ResponseWriter, req *http.Request) {
		var val string
		io.WriteString(w, "--------------------------------------\n")
		for rows.Next() {
			err := rows.Scan(clmvals...)
			if err != nil {
				fmt.Printf("Unable to Scan values from Select Query")
			}
			for i, col := range clmbox {
				val = string(col)
				io.WriteString(w, clm[i])
				io.WriteString(w, ":")
				io.WriteString(w, val)
				io.WriteString(w, "\n")
			}
			io.WriteString(w, "--------------------------------------\n")
		}
	}
	http.HandleFunc("/selectall", sqlhandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
