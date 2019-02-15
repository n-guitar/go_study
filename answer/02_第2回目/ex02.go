package main

// time.Sleepの引数にはtime.Durationという型を受け付ける。(time.Durationはint64の別名型)
// ナノ秒が基礎なので、time.Sleep(1)とやってしまうとたった1ナノ秒しか待たない。

import (
    "fmt"
    "os"
    "time"
    "bufio"
)

func main() {
const BUFSIZE = 1024 // 読み込みバッファのサイズ

    file, err := os.Open(`read.txt`) // <- fileをOpenする。
    if err != nil {
        fmt.Printf("Damn")
    }
    defer file.Close() // <- openしたfileをcloseする。os.Openと対で書かないといけない？

/*    buf := make([]byte, BUFSIZE)
    for {
        n, err := file.Read(buf)
        if n == 0 {
            break
        }
        if err != nil {
            // Readエラー処理
            break
        }

注意）コメントアウトしたところを含めると以下のように怒られる。
bufio.NewScannerの引数にタイプio.Readerとしてbuf（type [] byte）を使用することはできません：
[] byteはio.Readerを実装していません（Readメソッドがありません）

*/
        scanner := bufio.NewScanner(file)

        for scanner.Scan() {
            fmt.Println(scanner.Text())
            time.Sleep(1 * time.Second)
        }

        if err = scanner.Err(); err != nil {
        fmt.Printf("Damn")
        }
}
