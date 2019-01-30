# 以下をすべて実行せよ
- 内容  
    mysqlコンテナの起動  
    `port3306で外部からアクセス可能：127.0.0.1:3306`

※ docker-compose.yml存在するディレクトリ上で実行  
docker images作成
```bash
$ docker-compose build
```

docker コンテナ作成
```bash
$ docker-compose up -d

# 起動確認
$ docker-compose ps

# mysqlコンテナの起動確認
# mysql_003 | Version: '5.7.22'  socket: '/var/run/mysqld/mysqld.sock'  port: 3306  MySQL Community Server (GPL)と表示されればOK
$ docker-compose logs
```

sql実行
```bash
# db一覧
$ docker exec -it mysql_003 mysql -ugo_user -pgo_user go_apps -e 'show databases;'

# select文
docker exec -it mysql_003 mysql -ugo_user -pgo_user go_apps -e 'select * from tasks;'
```

コンテナ停止＆起動＆停止
```bash
# コンテナ停止
$ docker-compose stop; docker-compose ps

# コンテナ起動
$ docker-compose start; docker-compose ps

# コンテナ停止
$ docker-compose stop; docker-compose ps
```