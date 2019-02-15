package main

import (
     "fmt"
     "os"
)

func main() {

//    fmt.Printf("args count: %d\n", len(os.Args))
//    fmt.Printf("args : %#v\n", os.Args)

    for i, v := range os.Args {
        fmt.Printf("args[%d] -> %s\n", i, v)

    }
fmt.Printf("疑問：引数の数を限定しない時に、args[0]を表示しない方法を知りたい。\n")
}
