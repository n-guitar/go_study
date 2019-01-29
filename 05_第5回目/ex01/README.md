# 環境メモ


- go-sql-driver/mysql のインストール
```bash
go get github.com/go-sql-driver/mysql
```

- 参考ページ
https://golang.org/pkg/database/sql/

https://github.com/golang/go/wiki/SQLDrivers




- メモ
参考：https://www.mas9612.net/entry/go-mysql/   
SQLの実行は， DB.Exec() 及び DB.Query() メソッドを使う   
CREATE文やINSERT文など，DBからデータが返ってこないものに関しては DB.Exec() メソッドを用い   
SELECT文などDBからデータを取得するのが目的であるものに関しては DB.Query() メソッドを用いる．   
