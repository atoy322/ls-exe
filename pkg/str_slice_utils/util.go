package str_slice_utils

import (
    "regexp"
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
