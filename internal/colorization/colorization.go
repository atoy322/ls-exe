package colorization

import (
	"io/fs"
)

const (
    Reset  = "\x1b[0m"
    Red    = "\x1b[31m"
    Green  = "\x1b[32m"
    Yellow = "\x1b[33m"
    Blue   = "\x1b[34m"
    White  = "\x1b[37m"
    BgRed    = "\x1b[41m"
    BgGreen  = "\x1b[42m"
    BgYellow = "\x1b[43m"
    BgBlue   = "\x1b[44m"
    BgWhite  = "\x1b[47m"
)

type TextStyle struct {
    ForeColor string
    BgColor   string
    Bold      string
    Italic    string
}

func (ts *TextStyle) Colorize(text string) string {
    reset := ""
    if ts.ForeColor != "" || ts.BgColor != "" || ts.Bold != "" || ts.Italic != "" {
        reset = Reset
    }
    return ts.ForeColor + ts.BgColor + ts.Bold + ts.Italic + text + reset
}

func Judge(f fs.FileMode) *TextStyle {
    style := &TextStyle{}
    if f.IsDir() {
        style.ForeColor = Blue
    }
    return style
}
