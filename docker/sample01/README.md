# 以下をすべて実行せよ
※ Dockerfileが存在するディレクトリ上で実行

docker images作成
```bash
$ docker build .
```

mysqlコンテナ起動
※ --rm:コンテナ停止時に削除
```bash
$ docker run --rm --name mysql --hostname mysql -e MYSQL_ROOT_PASSWORD=root mysql:8.0
```

コンテナログイン
```bash
$ docker exec -it mysql /bin/bash
```

dbログイン
```bash
root@mysql:/# mysql -uroot -proot
```

sql操作
```sql
-- db一覧
mysql> show databases;

-- db作成
mysql> create database go_apps character set utf8mb4;
mysql> show databases;

-- user 作成
mysql> create user 'go_user'@'localhost' identified by 'go_user';
mysql> select user, host from mysql.user;

-- 権限の確認
mysql> show grants for 'go_user'@'localhost';

-- userにgo_apps dbのすべての権限を付与
mysql> grant all on go_apps.* to 'go_user'@'localhost';
mysql> flush privileges;
mysql> show grants for 'go_user'@'localhost';

-- go_userでログイン
mysql> exit
root@mysql:/# mysql -ugo_user -pgo_user
mysql> show databases;

-- db接続
mysql> use go_apps;

-- テーブル作成
mysql> create table tasks(id int auto_increment not null primary key, username varchar(255), task text, end_flag varchar(5), created_at timestamp not null default current_timestamp, updated_at timestamp not null default current_timestamp on update current_timestamp, index(id));

-- テーブル一覧
mysql> show tables;

-- データinsert
mysql> insert into tasks (username,task) values ("takita","go task");
mysql> insert into tasks (username,task) values ("nakahara","go task");

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
root@mysql:/# exit
$ docker stop mysql
```