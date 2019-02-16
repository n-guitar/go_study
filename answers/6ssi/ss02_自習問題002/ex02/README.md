# 以下をすべて実行せよ
- 内容  
    一時的なmysqlコンテナの起動  
    create系SQLの自動実行

※ Dockerfileが存在するディレクトリ上で実行  
docker images作成
```bash
$ docker build -t docker_sample:002 .
```

mysqlコンテナ起動
※ --rm:コンテナ停止時に削除
```bash
$ docker run --rm -d -v `pwd`/sql:/docker-entrypoint-initdb.d --name mysql_002 --hostname mysql_002 -e MYSQL_ROOT_PASSWORD=root docker_sample:002

# X Plugin ready for connections.と表示されれば起動
$ docker logs mysql_002
```

コンテナログイン
```bash
$ docker exec -it mysql_002 /bin/bash
```

sql操作
```sql
-- Goユーザでログイン
root@mysql_002:/# mysql -ugo_user -pgo_user

-- DB一覧
mysql> show databases;

-- db接続
mysql> use go_apps;

-- テーブル一覧
mysql> show tables;

-- テーブルselect
mysql> select * from tasks;

-- テーブルupdate
mysql> update tasks set end_flag="end" where id=1;
mysql> select * from tasks;
mysql> select * from tasks where end_flag="end";

-- 任意のデータ削除
mysql> delete from tasks where id=2;
mysql> select * from tasks;

-- テーブル削除
mysql> drop table tasks;
mysql> show tables;
```

コンテナ停止(削除)
```bash
mysql> exit
root@mysql_002:/# exit
$ docker stop mysql_002
```