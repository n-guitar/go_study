package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	id         int
	username   string
	task       string
	end_flag   sql.NullString
	created_at string
	update_at  string
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:9999", nil))
}

func tag(tag string, any interface{}) string {

	res := ""
	switch v := any.(type) {
	case string:
		res = "<" + tag + ">" + v + "</" + tag + ">"
	case sql.NullString:
		res = "<" + tag + ">" + v.String + "</" + tag + ">"
	default:
		fmt.Println(reflect.TypeOf(any))
		res = "<" + tag + ">" + "</" + tag + ">"
	}

	return res
}

func pretty_html(columns []string, users []User) string {

	var th string
	for _, column := range columns {
		th += tag("th", column)
	}
	th = tag("tr", th)

	var td string
	for _, u := range users {

		tmp_td := []string{
			tag("td", strconv.Itoa(u.id)),
			tag("td", u.username),
			tag("td", u.task),
			tag("td", u.end_flag),
			tag("td", u.created_at),
			tag("td", u.update_at)}

		td += tag("tr", strings.Join(tmp_td, ""))
	}

	res := tag("tbody", th+td)
	res = tag("table", res)

	if len(users) == 0 {
		res += tag("p", "sorry.. record not found.")
	}

	return res
}

func handler(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	if method == "POST" {
		postHandler(w, r)
	}
	getHandler(w, r)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	id := r.FormValue("no")

	db, err := sql.Open("mysql", "go_user:go_user@/go_apps")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	sql := "update tasks set end_flag = 'end' where id = " + id
	if id == "" {
		return
	}

	_, err = db.Exec(sql)
	if err != nil {
		panic(err.Error())
	}

}

func getHandler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	id := r.FormValue("no")

	db, err := sql.Open("mysql", "go_user:go_user@/go_apps")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	sql := "SELECT * FROM tasks"
	if id != "" {
		sql += " where id = " + id
	}
	rows, err := db.Query(sql)
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
