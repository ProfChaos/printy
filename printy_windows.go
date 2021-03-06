// Copyright 2015 Dan S. Audne. All rights reserved.

package printy

import (
	"fmt"
	"os"
	"syscall"
	"time"

	"github.com/azer/is-terminal"
)

var (
	white    = "\033[37m"
	reset    = "\033[0m"
	red      = "\033[31m"
	blue     = "\033[34m"
	yellow   = "\033[33m"
	bgRed    = "\033[41m"
	bgYellow = "\033[43m"
	bgBlue   = "\033[44m"
)

type style struct {
	Time string
	Type string
	Text string
}

// Err for logging error messages on stderr
func Err(s interface{}) {
	f := style{
		Time: red,
		Type: bgRed + white,
	}
	format(os.Stderr, syscall.Stderr, s, "Error", f)
}

// Warn is for logging warnings on stdout
func Warn(s interface{}) {
	f := style{
		Time: yellow,
	}
	format(os.Stdout, syscall.Stdout, s, "Warning", f)
}

// Info for marking more important log messages stdout
func Info(s interface{}) {
	f := style{
		Time: blue,
	}
	format(os.Stdout, syscall.Stdout, s, "Info", f)
}

// Log is log
func Log(s interface{}) {
	format(os.Stdout, syscall.Stdout, s, "Log", style{})
}

// format formats output
func format(f *os.File, sys syscall.Handle, s interface{}, typ string, c style) {
	now := time.Now().Format("2006-01-02 15:04:05.000")
	r := reset

	if !isterminal.IsTerminal(int(sys)) {
		c = style{}
		r = ""
	}

	fmt.Fprintf(f, "%v[%v] %v%v:%v %v%v\n", c.Time, now, c.Type, typ, r, s, r)
}
