package pretty_print

import (
    "fmt"
    "regexp"

    "github.com/nathan-fiscaletti/consolesize-go"
)

func MaxLen(str_slice []string) int {
    re, _ := regexp.Compile("\x1b.[0-9]{1,2}m")
    max := 0
    var str string
    for _, i := range str_slice {
        str = re.ReplaceAllString(i, "")
        if max < len(str) {
            max = len(str)
        }
    }
    return max
}

func Pprint(filenames []string) {
    w, _ := consolesize.GetConsoleSize()
    max := MaxLen(filenames)
 
    col := w / max
    for i, filename := range filenames {
        fmt.Printf("%*s", max, filename)

        if (i+1) % col == 0 {
            fmt.Print("\n")
        }
    }
}
