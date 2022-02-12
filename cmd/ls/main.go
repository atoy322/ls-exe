package main

import (
	"fmt"
	"io/fs"
	"os"

	"github.com/atoy322/ls-exe/internal/cmdline_parser"
	"github.com/atoy322/ls-exe/internal/colorization"
	"github.com/atoy322/ls-exe/internal/pretty_print"
)

func main() {
    options := cmdline_parser.Parse(os.Args)
    // options.Color = true
    finfo, e := os.Stat(options.Directory)

    ls := make([]os.DirEntry, 0)
    if e != nil {
        fmt.Println("ls: No such file or directory.")
    } else if !finfo.IsDir() {  // is file
        ls = append(ls, fs.FileInfoToDirEntry(finfo))
    } else {  // is folder
        ls, e = os.ReadDir(options.Directory)

        if e != nil {
            panic(e)
        }
    }

    var style colorization.TextStyle
    filenames := make([]string, 0)
    for _, i := range ls {
        style = colorization.Judge(i.Type())

        if i.Name()[0] == '.' && !options.All {
            continue
        }
        if options.Color {
            filenames = append(filenames, style.Colorize(i.Name()))
        } else {
            filenames = append(filenames, i.Name())
        }
    }

    pretty_print.Pprint(filenames)
    fmt.Print("\n")
}
