package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
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
		var nameslc []string
		var passslc []string
		query := r.URL.Query()
		for _, rname := range query["name"] {
			nameslc = append(nameslc, rname)
		}
		for _, rpass := range query["password"] {
			passslc = append(passslc, rpass)
		}
		fmt.Printf("%d name POSTed -> %s \n", len(nameslc), query["name"])
		fmt.Printf("%d password POSTed -> %s \n", len(passslc), query["password"])
		if len(nameslc) != len(passslc) {
			io.WriteString(w, "[ERROR]Either name or password is missing.\n")
		} else {
			for i := range nameslc {
				_, err := db.Exec("INSERT INTO users(name,password) VALUES(?,?)", nameslc[i], passslc[i])
				if err != nil {
					io.WriteString(w, "[ERROR]Failed to INSERT INTO table.\n")
				}
			}
			io.WriteString(w, "[SUCCESS]Your POSTed data was INSERTed INTO table.\n")
		}
	} else {
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
		io.WriteString(w, "-----------------------------------\n")
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
}
