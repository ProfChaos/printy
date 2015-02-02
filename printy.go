// Copyright 2015 Dan S. Audne. All rights reserved.

package printy

import (
	"fmt"
	"github.com/azer/is-terminal"
	"os"
	"syscall"
	"time"
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
func format(f *os.File, sys int, s interface{}, typ string, c style) {
	now := time.Now().Format("2006-02-01 03:04:05.000")
	r := reset

	if !isterminal.IsTerminal(sys) {
		c = style{}
		r = ""
	}

	fmt.Fprintf(f, "%s[%s] %s%s:%s %s%s\n", c.Time, now, c.Type, typ, r, s, r)
}
