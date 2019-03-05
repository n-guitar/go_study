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
	sqlhandler := func(w http.ResponseWriter, req *http.Request) {
		var val string
		var fndno []string
		//var nfndno []string
		query := req.URL.Query()
		if query["no"] == nil {
			io.WriteString(w, "\"No\" not specified.\n")
		} else {
			for _, idnum := range query["no"] {
				rows, err := db.Query("SELECT * FROM tasks where id=?", idnum)
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

				for rows.Next() {
					err := rows.Scan(clmvals...)
					fmt.Println("Query's ID is " + idnum)
					fmt.Println("Result ID is " + string(clmbox[0]))
					fmt.Println("Result Name is " + string(clmbox[1]))
					fndno = append(fndno, idnum)
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

			all := make([]string, len(query["no"]))

			for a, noval := range query["no"] {
				all[a] = noval
			}
			for _, fn := range fndno {
				for a, qno := range query["no"] {
					if fn == qno {
						all[a] = "OK"
					}
				}
			}
			for _, n := range all {
				if n != "OK" {
					io.WriteString(w, "ID "+n+" are not found\n")
				}
			}
		}
	}
	http.HandleFunc("/tasks", sqlhandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
