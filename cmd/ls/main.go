package main

import (
    "fmt"
    "os"

    "github.com/atoy322/ls-exe/internal/colorization"
)

func main() {
    ls, e := os.ReadDir(".")
    if e != nil {
        panic(e)
    }
    style := &colorization.TextStyle{}
    for _, i := range ls {
        style = colorization.Judge(i.Type())
        fmt.Println(style.Colorize(i.Name()))
    }
}
