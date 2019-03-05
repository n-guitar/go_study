package main

import (
	"crypto/sha256"
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
		var nameslc []string
		var rnameslc []string
		var passslc []string
		var hashslc []string
		query := r.URL.Query()
		for _, rname := range query["name"] {
			nameslc = append(nameslc, rname)
			rnameslc = append(rnameslc, rname)
		}
		for _, rpass := range query["password"] {
			passslc = append(passslc, rpass)
		}

		for h, rname := range nameslc {
			hashslc = append(hashslc, fmt.Sprintf("%x", sha256.Sum256([]byte(rname+passslc[h]))))
		}

		fmt.Printf("%d name POSTed -> %s \n", len(nameslc), query["name"])
		fmt.Printf("%d password POSTed -> %s \n", len(passslc), query["password"])
		if len(nameslc) != len(passslc) {
			io.WriteString(w, "[ERROR]Either name or password is missing.\n")
		} else {
			for i := range nameslc {
				rows, err := db.Query("SELECT name,password,hash FROM users WHERE name = ?", nameslc[i])
				if err != nil {
					io.WriteString(w, "[ERROR] Failed to check the table.\n")
				}
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
				for v := range values {
					scanArgs[v] = &values[v]
				}
				for rows.Next() {
					err = rows.Scan(scanArgs...)
					if err != nil {
						log.Fatalln(err)
					}

					for c, col := range values {
						fmt.Printf("column : %s , value : %s\n", columns[c], string(col))
						if columns[c] == "hash" {
							if hashslc[i] == string(col) {
								nameslc[i] = fmt.Sprintf("%x", sha256.Sum256([]byte(nameslc[i]+passslc[i])))
							}
						}
					}
				}
				if nameslc[i] == rnameslc[i] {
					io.WriteString(w, "NG\n")
				} else {
					io.WriteString(w, nameslc[i]+"\n")
				}
			}
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
