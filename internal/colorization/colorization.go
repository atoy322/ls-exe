package colorization

import (
    "strings"
	"io/fs"
)

const (
    Reset    = "\x1b[0m"
    Red      = "\x1b[31m"
    Green    = "\x1b[32m"
    Yellow   = "\x1b[33m"
    Blue     = "\x1b[34m"
    Magenta  = "\x1b[35m"
    Cyan     = "\x1b[36m"
    White    = "\x1b[37m"
    BgRed    = "\x1b[41m"
    BgGreen  = "\x1b[42m"
    BgYellow = "\x1b[43m"
    BgBlue   = "\x1b[44m"
    BgMagenta= "\x1b[45m"
    BgCyan   = "\x1b[46m"
    BgWhite  = "\x1b[47m"
    Bold     = "\x1b[1m"
    Italic   = "\x1b[3m"
    Undrline = "\x1b[4m"
    Cancel   = "\x1b[9m"
)

type TextStyle struct {
    ForeColor string
    BgColor   string
    Bold      bool
    Italic    bool
    Undrline  bool
    Cancel    bool
}

func (ts TextStyle) Colorize(text string) string {
    others := ""
    if ts.Bold     { others += Bold     }
    if ts.Italic   { others += Italic   }
    if ts.Undrline { others += Undrline }
    if ts.Cancel   { others += Cancel   }

    reset := ""
    if ts.ForeColor != "" || ts.BgColor != "" || others != "" {
        reset = Reset
    }
    return ts.ForeColor + ts.BgColor + others + text + reset
}

func Judge(f fs.DirEntry) TextStyle {
    var style TextStyle

    if f.IsDir() {   // is dir

        style.ForeColor = Blue
        style.Bold = true

    } else {         // is file

        cs := strings.Split(f.Name(), ".")
        // name := strings.Join(cs[:len(cs)-1], ".")
        ext := cs[len(cs)-1]

        if ext == "exe" {    // is executable

            style.ForeColor = Green
            style.Bold = true

        } else if ext == "py" || ext == "c" || ext == "go" || ext == "sh" || ext == "js" || ext == "bat" {
            style.Undrline = true
        }

    }

    return style
}
