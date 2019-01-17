# goqueryインストール
# goqueryインストール
https://github.com/PuerkitoBio/goquery

```bash
# インストール
$ go get github.com/PuerkitoBio/goquery

# テスト
$ export $OPATH=`go env | grep GOPATH | sed 's/^.*"\(.*\)".*$/\1/'`
$ cd $GOPATH/src/github.com/PuerkitoBio/goquery
$ go test
```