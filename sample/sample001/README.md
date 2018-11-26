# WEBサーバの例
サーバへアクセスしたときに指定するURLのパスを返すWEBサーバ  
例えば、http://localhost:8000/hello   に対してアクセスするとURL.Path = "/hello"となる。

> http.HandleFunc   

URLを受け取り、func(ResponseWriter, *Request)を渡す

> handler(w http.ResponseWriter, r *http.Request)   

リクエストされたURL rのPath要素を返す
実際にはfunc(ResponseWriter, *Request)のこと
http.Handlerにキャストする

> http.Handler   

HTTPリクエストを受けてレスポンスを返す
ServeHTTP関数をもつインターフェース
```go
type Handler interface {
  ServeHTTP(ResponseWriter, *Request)
}
```

 
 # 参考
## httpパッケージについて   
 http://golang.jp/pkg/http

## キャスト   
 https://wa3.i-3-i.info/word1190.html    
 置き換えると思っっておけば良い（多分）

## インターフェース   
 ここでは詳しく説明しない
 インターフェースという概念を理解するにはもう少しあとにしようと思う。
 Javaとかのクラスの継承がわかる人なら、同じような感じと思っておいてほしい。

 intとかstringとかの言わるる型は想像つくと思う
 インターフェースとは方の一種であり、メソッドを定義する
 インターフェースが定義された型が、どのようなメソッドを持つべきかを定義できるもので
 なんかよくわからないけど、便利なものと思っておいてほしい。




# code 
```go
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
```