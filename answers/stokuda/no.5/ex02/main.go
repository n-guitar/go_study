package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

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
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:9999", nil))
}

func tag(tag string, value string) string {
	return "<" + tag + ">" + value + "</" + tag + ">"
}

func pretty_html(columns []string, users []User) string {

	var th string
	for _, column := range columns {
		th += tag("th", column)
	}
	th = tag("tr", th)

	var td string
	for _, u := range users {

		end_flag, ok := u.end_flag.(string)
		if !ok {
			end_flag = "null"
		}

		tmp_td := []string{
			tag("td", strconv.Itoa(u.id)),
			tag("td", u.username),
			tag("td", u.task),
			tag("td", end_flag),
			tag("td", u.created_at),
			tag("td", u.update_at)}

		td += tag("tr", strings.Join(tmp_td, ""))
	}

	res := tag("tbody", th+td)
	res = tag("table", res)
	return res
}

func handler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

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

	var user User
	var users []User
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
		users = append(users, user)
	}

	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "%s", pretty_html(columns, users))

}
