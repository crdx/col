package col

import (
	"os"
)

var enable = true

// Enable enables colours.
func Enable() {
	enable = true
}

// Disable disables colours.
func Disable() {
	enable = false
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
func Green(str string) string { return render(str, fgGreen, false) }

// Blue sets the foreground colour to blue.
func Blue(str string) string { return render(str, fgBlue, false) }

// Red sets the foreground colour to red.
func Red(str string) string { return render(str, fgRed, false) }

// Yellow sets the foreground colour to yellow.
func Yellow(str string) string { return render(str, fgYellow, false) }

// Magenta sets the foreground colour to magenta.
func Magenta(str string) string { return render(str, fgMagenta, false) }

// White sets the foreground colour to white.
func White(str string) string { return render(str, fgWhite, false) }

// Black sets the foreground colour to black.
func Black(str string) string { return render(str, fgBlack, false) }

// Cyan sets the foreground colour to cyan.
func Cyan(str string) string { return render(str, fgCyan, false) }

// BgGreen sets the background colour to green.
func BgGreen(str string) string { return render(str, bgGreen, true) }

// BgBlue sets the background colour to blue.
func BgBlue(str string) string { return render(str, bgBlue, true) }

// BgRed sets the background colour to red.
func BgRed(str string) string { return render(str, bgRed, true) }

// BgYellow sets the background colour to yellow.
func BgYellow(str string) string { return render(str, bgYellow, true) }

// BgMagenta sets the background colour to magenta.
func BgMagenta(str string) string { return render(str, bgMagenta, true) }

// BgWhite sets the background colour to white.
func BgWhite(str string) string { return render(str, bgWhite, true) }

// BgBlack sets the background colour to black.
func BgBlack(str string) string { return render(str, bgBlack, true) }

// BgCyan sets the background colour to cyan.
func BgCyan(str string) string { return render(str, bgCyan, true) }

// Underline sets the styling to underlined.
func Underline(str string) string { return render(str, underline, false) }

// Bold sets the styling to bold.
func Bold(str string) string { return render(str, bold, false) }

// Italic sets the styling to italic.
func Italic(str string) string { return render(str, italic, false) }

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

func render(str string, code string, bg bool) string {
	if !enable {
		return str
	}

	if bg {
		code = ";" + code
	}

	return escape + code + "m" + str + escape + reset + "m"
}
