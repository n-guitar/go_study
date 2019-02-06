/*
初期作成パターン
nullが存在するがSQL文でカバー(COALESCE)
db connectionを分割
*/
package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// ToDo 排他ロック制御追加

type Tasks struct {
	id         int
	username   string
	task       string
	end_flag   string
	created_at string
	updated_at string
}

// db connection
func conDatabase() *sql.DB {
	var db *sql.DB
	db, err := sql.Open("mysql", "go_user:go_user@tcp(localhost:3306)/go_apps")
	if err != nil {
		log.Fatalf("sql.Open(): %s\n", err)
	}
	return db
}

func main() {

	http.HandleFunc("/taskall", func(w http.ResponseWriter, r *http.Request) {

	})
	err := http.ListenAndServe("localhost:8000", nil)

}

func selectTable(db *sql.DB) *Tasks {
	db = conDatabase()

	// SQL文
	rows, err := db.Query("SELECT id,username,task,COALESCE(end_flag,'null'),created_at,updated_at FROM tasks")
	if err != nil {
		log.Fatalf("db.Query(): %s\n", err)
	}

	for rows.Next() {
		m := Tasks{}
		err = rows.Scan(&m.id, &m.username, &m.task, &m.end_flag, &m.created_at, &m.updated_at)
		if err != nil {
			log.Fatalf("rows.Scan(): %s\n", err)
		}

		// fmt.Fprintf(w, "id: %d,\t username: %s,\t task: %s,\t end_flag: %s,\t created_at: %s,\t updated_at: %s\n", m.id, m.username, m.task, m.end_flag, m.created_at, m.updated_at)

		if err = rows.Err(); err != nil {
			log.Fatalf("rows.Err(): %s\n", err)
		}
		defer db.Close()
		return m
	}

	// func newPoint3d(x, y, z float64) *Point3d {
	// 	p := new(Point3d)
	// 	p.x, p.y, p.z = x, y, z
	// 	p.distance = func(q *Point3d) float64 { return distance3d(p, q) }
	// 	return p
	// }

	// type Point struct {
	// 	x, y float64
	// 	distance func(*Point) float64
	// }

}
