package main

import (
	"flag"
	"fmt"
)

func main() {
	// パラメータ名は,パラメータ名、デフォルト値、説明文
	//　戻り値はポインタ
	i := flag.Int("flag1", 100, "Int型の整数を指定")
	s := flag.String("flag2", "con-majikayo", "String型の文字列を指定")

	// パラメータの解析
	flag.Parse()

	// 取得したパラメータの出力
	fmt.Println(*i, *s)

}
