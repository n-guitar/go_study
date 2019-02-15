package main

import (
    "fmt"
    "os"
//    "io/ioutil"
)

func main() {

/*
    //　読み込みファイルオープン
    file,err := os.Open(`read.txt`)
    // Openエラー処理
    if err != nil {
       fmt.Println(err)
        return
    }
    fmt.Println(file)  // os.Openはファイルをオープンするもの？で、Read関数でByte配列へ渡す？
                       // じゃないと文字として表示できない。。。
}
*/

const BUFSIZE = 1024 // 読み込みバッファのサイズ

    file, err := os.Open(`read.txt`) // <- fileをOpenする。
    if err != nil {
        // Openエラー処理
    }
    defer file.Close() // <- openしたfileをcloseする。os.Openと対で書かないといけない？

    buf := make([]byte, BUFSIZE)
    for {
        n, err := file.Read(buf)
        if n == 0 {
            break
        }
        if err != nil {
            // Readエラー処理
            break
        }

        fmt.Print(string(buf[:n]))
    }
}



/* os.Openを使わないものもあるっぽい・・・
func main() {
    data, err := ioutil.ReadFile(`read.txt`)
    if err != nil {
        // エラー処理
    }
    fmt.Print(string(data))
}
*/
