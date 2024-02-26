package col

import (
	"fmt"
	"os"
)

var enabled = true

// Enable enables colours.
func Enable() {
	enabled = true
}

// Disable disables colours.
func Disable() {
	enabled = false
}

// Init initialises colours based on heuristics.
//
// The following heuristics are used to figure out whether colour support should be disabled.
//
//   - The NO_COLOR environment variable is set.
//   - The terminal is not interactive.
func Init() {
	InitUnless(false)
}

// InitUnless initialises colours based on heuristics unless overridden.
//
// Like [Init], but optionally override the heuristics and force colours to be disabled if disable
// is set to true.
func InitUnless(disable bool) {
	if disable {
		Disable()
		return
	}

	if os.Getenv("NO_COLOR") != "" {
		Disable()
		return
	}

	if !isInteractive() {
		Disable()
		return
	}
}

// Green sets the foreground colour to green.
func Green(fmt any, args ...any) string { return render(fmt, fgGreen, false, args...) }

// Blue sets the foreground colour to blue.
func Blue(fmt any, args ...any) string { return render(fmt, fgBlue, false, args...) }

// Red sets the foreground colour to red.
func Red(fmt any, args ...any) string { return render(fmt, fgRed, false, args...) }

// Yellow sets the foreground colour to yellow.
func Yellow(fmt any, args ...any) string { return render(fmt, fgYellow, false, args...) }

// Magenta sets the foreground colour to magenta.
func Magenta(fmt any, args ...any) string { return render(fmt, fgMagenta, false, args...) }

// White sets the foreground colour to white.
func White(fmt any, args ...any) string { return render(fmt, fgWhite, false, args...) }

// Black sets the foreground colour to black.
func Black(fmt any, args ...any) string { return render(fmt, fgBlack, false, args...) }

// Cyan sets the foreground colour to cyan.
func Cyan(fmt any, args ...any) string { return render(fmt, fgCyan, false, args...) }

// BgGreen sets the background colour to green.
func BgGreen(fmt any, args ...any) string { return render(fmt, bgGreen, true, args...) }

// BgBlue sets the background colour to blue.
func BgBlue(fmt any, args ...any) string { return render(fmt, bgBlue, true, args...) }

// BgRed sets the background colour to red.
func BgRed(fmt any, args ...any) string { return render(fmt, bgRed, true, args...) }

// BgYellow sets the background colour to yellow.
func BgYellow(fmt any, args ...any) string { return render(fmt, bgYellow, true, args...) }

// BgMagenta sets the background colour to magenta.
func BgMagenta(fmt any, args ...any) string { return render(fmt, bgMagenta, true, args...) }

// BgWhite sets the background colour to white.
func BgWhite(fmt any, args ...any) string { return render(fmt, bgWhite, true, args...) }

// BgBlack sets the background colour to black.
func BgBlack(fmt any, args ...any) string { return render(fmt, bgBlack, true, args...) }

// BgCyan sets the background colour to cyan.
func BgCyan(fmt any, args ...any) string { return render(fmt, bgCyan, true, args...) }

// Underline sets the styling to underlined.
func Underline(fmt any, args ...any) string { return render(fmt, underline, false, args...) }

// Bold sets the styling to bold.
func Bold(fmt any, args ...any) string { return render(fmt, bold, false, args...) }

// Italic sets the styling to italic.
func Italic(fmt any, args ...any) string { return render(fmt, italic, false, args...) }

// —————————————————————————————————————————————————————————————————————————————————————————————————

func isInteractive() bool {
	fileInfo, _ := os.Stdout.Stat()
	return (fileInfo.Mode() & os.ModeCharDevice) != 0
}

const (
	escape = "\033["

	fgBlack   = "30"
	fgRed     = "31"
	fgGreen   = "32"
	fgYellow  = "33"
	fgBlue    = "34"
	fgMagenta = "35"
	fgCyan    = "36"
	fgWhite   = "37"

	bgBlack   = "40"
	bgRed     = "41"
	bgGreen   = "42"
	bgYellow  = "43"
	bgBlue    = "44"
	bgMagenta = "45"
	bgCyan    = "46"
	bgWhite   = "47"

	underline = "4"
	italic    = "3"
	bold      = "1"
	reset     = "0"
)

func render(v any, code string, bg bool, args ...any) string {
	s := fmt.Sprint(v)

	if len(args) > 0 {
		s = fmt.Sprintf(s, args...)
	}

	if !enabled {
		return s
	}

	if bg {
		code = ";" + code
	}

	return escape + code + "m" + s + escape + reset + "m"
}
