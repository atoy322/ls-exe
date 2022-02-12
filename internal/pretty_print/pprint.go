package pretty_print

import (
    "fmt"
)

func Pprint(filenames []string) {
    for _, filename := range filenames {
        fmt.Printf("%10s\t", filename)
    }
}
