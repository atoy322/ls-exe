package str_slice_utils

import (
    "regexp"
)


func MaxLen(str_slice []string) int {
    max := 0
    
    for _, i := range str_slice {
        if max < len(i) {
            max = len(i)
        }
    }

    return max
}

func Filter(str_slice []string, from, to string) ([]string, error) {
    re, e := regexp.Compile(from)
    if e != nil {
        return []string{}, e
    }

    var str string
    result := make([]string, 0)
    for _, i := range str_slice {
        str = re.ReplaceAllString(i, to)
        result = append(result, str)
    }
    return result, nil
}
