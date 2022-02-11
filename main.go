package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/nathan-fiscaletti/consolesize-go"
)

type FileInfoToShow struct {
    Color string
    BgColor string
    Name string
}

func (fis *FileInfoToShow) ToColorExp() string {
    switch fis.Color {
        case "red":
            return "\x1b[31m"
        default:
            return ""
    }
}

func main() {
    w, h := consolesize.GetConsoleSize()
    fmt.Println(w, h)
    var all bool
    var color bool
    for i, command := range os.Args {
        fmt.Printf("Cmd[%d] = %s\n", i, command)
        if command[0] == '-' {
            if strings.Contains(command, "a") {
                all = true
            }
            if strings.Contains(command, "color") {
                color = true
            }
        }
    }
    files, e := os.ReadDir(".")
    if e != nil {
        panic(e)
    }

    var isreg bool
    var color_expression string

    for _, entry := range files {
        isreg = entry.Type().IsRegular()
        if entry.IsDir() && color {
            color_expression = "\x1b[34m"
        } else {
            color_expression = ""
        }
        if isreg || all {
            fmt.Println(color_expression+entry.Name()+"\x1b[0m")
        }
    }
}
