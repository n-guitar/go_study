/* 初期作成パターン
 */
package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// ToDo 排他ロック制御追加

type Tasks struct {
	id         int
	username   string
	task       string
	end_flag   sql.NullString // null判定
	created_at string
	updated_at string
}

// ToDo 関数の分割 SQL実行と接続部分で分ける

func main() {
	// db, err := sql.Open("mysql", "user:password@/dbname")
	var db *sql.DB
	db, err := sql.Open("mysql", "go_user:go_user@tcp(localhost:3306)/go_apps")
	if err != nil {
		log.Fatalf("sql.Open(): %s\n", err)
	}
	defer db.Close()

	// SQL文
	rows, err := db.Query("SELECT id,username,task,end_flag,created_at,updated_at FROM tasks")
	if err != nil {
		log.Fatalf("db.Query(): %s\n", err)
	}

	for rows.Next() {
		m := Tasks{}
		// ToDo nullの場合の条件分岐を追加し、nullの場合はnullと表示する
		err = rows.Scan(&m.id, &m.username, &m.task, &m.end_flag, &m.created_at, &m.updated_at)
		if err != nil {
			log.Fatalf("rows.Scan(): %s\n", err)
		}

		fmt.Printf("id: %d,\t username: %s,\t task: %s,\t end_flag: %s,\t created_at: %s,\t updated_at: %s\n", m.id, m.username, m.task, m.end_flag, m.created_at, m.updated_at)

		if err = rows.Err(); err != nil {
			log.Fatalf("rows.Err(): %s\n", err)
		}
	}

}
