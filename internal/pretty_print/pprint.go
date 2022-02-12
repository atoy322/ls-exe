package pretty_print

import (
    "fmt"
    "strings"

    "github.com/nathan-fiscaletti/consolesize-go"
    "github.com/atoy322/ls-exe/pkg/str_slice_utils"
)


func Pprint(filenames []string) {
    w, _ := consolesize.GetConsoleSize()  // It becomes 0 when using pipe output. (Ex. "ls | grep *.go")

    filenames_without_color, _ := str_slice_utils.Filter(filenames, "\x1b.[0-9]{1,2}m", "")
    max := str_slice_utils.MaxLen(filenames_without_color)
    max += 3  // space between longest filename and other.

    col := w / max
    if col == 0 {  // When using pipe output.
        col = 1  // To prevent division by zero error.
    }
    var space_length int

    for i, filename := range filenames {
        space_length = max-len(filenames_without_color[i])
        fmt.Print(filename + strings.Repeat(" ", space_length))

        if (i+1) % col == 0 {
            fmt.Print("\n")
        }
    }
}
