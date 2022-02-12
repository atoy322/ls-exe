package cmdline_parser

import (
    "strings"
)

type RunOptions struct {
    All       bool
    Color     bool
    Directory string
}

func Parse(args []string) *RunOptions {
    options := &RunOptions{
        All: false,
        Color: false,
        Directory: "",
    }
    for _, arg := range args[1:] {
        if arg[0] == '-' {
            if strings.Contains(arg, "a") {
                options.All = true
            }
            if strings.Contains(arg, "color") {
                options.Color = true
            }
        } else {
            options.Directory = arg
        }
    }
    if options.Directory == "" {
        options.Directory = "."
    }
    return options
}
