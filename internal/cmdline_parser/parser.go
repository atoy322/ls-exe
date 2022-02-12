package cmdline_parser

import (
    "strings"
)

type RunOptions struct {
    All bool
    Color bool
    Directory string
}

func Parse(args []string) *RunOptions {
    options := &RunOptions{}
    for _, arg := range args {
        if strings.Contains(arg, "a") {
            options.All = true
        }
    }
    return options
}
