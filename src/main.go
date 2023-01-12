package col

import (
	"os"
)

var enable = true

func Enable()  { enable = true }
func Disable() { enable = false }

func Init() {
	InitUnless(false)
}

// Initialise colours.
// Use the `disable` parameter to override autodetection.
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

func Green(str string) string   { return render(str, fgGreen, false) }
func Blue(str string) string    { return render(str, fgBlue, false) }
func Red(str string) string     { return render(str, fgRed, false) }
func Yellow(str string) string  { return render(str, fgYellow, false) }
func Magenta(str string) string { return render(str, fgMagenta, false) }
func White(str string) string   { return render(str, fgWhite, false) }
func Black(str string) string   { return render(str, fgBlack, false) }
func Cyan(str string) string    { return render(str, fgCyan, false) }

func BgGreen(str string) string   { return render(str, bgGreen, true) }
func BgBlue(str string) string    { return render(str, bgBlue, true) }
func BgRed(str string) string     { return render(str, bgRed, true) }
func BgYellow(str string) string  { return render(str, bgYellow, true) }
func BgMagenta(str string) string { return render(str, bgMagenta, true) }
func BgWhite(str string) string   { return render(str, bgWhite, true) }
func BgBlack(str string) string   { return render(str, bgBlack, true) }
func BgCyan(str string) string    { return render(str, bgCyan, true) }

func Underline(str string) string { return render(str, underline, false) }
func Bold(str string) string      { return render(str, bold, false) }
func Italic(str string) string    { return render(str, italic, false) }
