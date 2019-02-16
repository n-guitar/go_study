package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"log"
	"net/http"
)

var db *sql.DB

func init() {
	const (
		dbName = "db01"
		dbUser = "app01"
		dbPass = "app01"
		dbHost = "127.0.0.1"
		dbPort = "3306"
	)

	dbSource := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8",
		dbUser,
		dbPass,
		dbHost,
		dbPort,
		dbName)

	var err error
	db, err = sql.Open("mysql", dbSource)
	if err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:9999", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	if method == "POST" {

	}
	io.WriteString(w, "test")
	rows, err := db.Query("select * from users")
	defer rows.Close()
	if err != nil {
		panic(err.Error())
	}

	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}
	values := make([]sql.RawBytes, len(columns))
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
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			io.WriteString(w, columns[i]+":"+value+"\n")
		}
		io.WriteString(w, "-----------------------------------\n")
	}
}
