package col

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnableDisable(t *testing.T) {
	Enable()
	assert.True(t, enabled)

	Disable()
	assert.False(t, enabled)

	Enable()
	assert.True(t, enabled)
}

func TestInitUnless(t *testing.T) {
	Enable()
	InitUnless(true)
	assert.False(t, enabled)
	Enable()

	os.Setenv("NO_COLOR", "1")
	InitUnless(false)
	assert.False(t, enabled)
	os.Unsetenv("NO_COLOR")
	Enable()
}

func TestInit(t *testing.T) {
	Enable()
	os.Setenv("NO_COLOR", "1")
	Init()
	assert.False(t, enabled)
	os.Unsetenv("NO_COLOR")
	Enable()
}

func TestColoursEnabled(t *testing.T) {
	Enable()

	tests := []struct {
		name     string
		fn       func(any, ...any) string
		expected string
	}{
		// Foreground
		{"Green", Green, "\033[32mtest\033[0m"},
		{"Blue", Blue, "\033[34mtest\033[0m"},
		{"Red", Red, "\033[31mtest\033[0m"},
		{"Yellow", Yellow, "\033[33mtest\033[0m"},
		{"Magenta", Magenta, "\033[35mtest\033[0m"},
		{"White", White, "\033[37mtest\033[0m"},
		{"Black", Black, "\033[30mtest\033[0m"},
		{"Cyan", Cyan, "\033[36mtest\033[0m"},
		// Background
		{"BgGreen", BgGreen, "\033[;42mtest\033[0m"},
		{"BgBlue", BgBlue, "\033[;44mtest\033[0m"},
		{"BgRed", BgRed, "\033[;41mtest\033[0m"},
		{"BgYellow", BgYellow, "\033[;43mtest\033[0m"},
		{"BgMagenta", BgMagenta, "\033[;45mtest\033[0m"},
		{"BgWhite", BgWhite, "\033[;47mtest\033[0m"},
		{"BgBlack", BgBlack, "\033[;40mtest\033[0m"},
		{"BgCyan", BgCyan, "\033[;46mtest\033[0m"},
		// Styles
		{"Bold", Bold, "\033[1mtest\033[0m"},
		{"Dim", Dim, "\033[2mtest\033[0m"},
		{"Italic", Italic, "\033[3mtest\033[0m"},
		{"Underline", Underline, "\033[4mtest\033[0m"},
		{"Strikethrough", Strikethrough, "\033[9mtest\033[0m"},
		{"Reverse", Reverse, "\033[7mtest\033[0m"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.fn("test"))
		})
	}
}

func TestColoursDisabled(t *testing.T) {
	Disable()
	defer Enable()

	fns := []func(any, ...any) string{
		Green, Blue, Red, Yellow, Magenta, White, Black, Cyan,
		BgGreen, BgBlue, BgRed, BgYellow, BgMagenta, BgWhite, BgBlack, BgCyan,
		Bold, Dim, Italic, Underline, Strikethrough, Reverse,
	}

	for _, fn := range fns {
		assert.Equal(t, "test", fn("test"))
	}
}

func TestFormatStringArgs(t *testing.T) {
	Enable()
	assert.Equal(t, "\033[32mhello world\033[0m", Green("hello %s", "world"))
	assert.Equal(t, "\033[31m5 errors\033[0m", Red("%d errors", 5))
	assert.Equal(t, "\033[1mcount: 42\033[0m", Bold("%s: %d", "count", 42))

	Disable()
	assert.Equal(t, "hello world", Green("hello %s", "world"))
	Enable()
}

func TestNonStringInput(t *testing.T) {
	Enable()
	assert.Equal(t, "\033[31m123\033[0m", Red(123))
	assert.Equal(t, "\033[34m45.67\033[0m", Blue(45.67))
}

func TestEmptyString(t *testing.T) {
	Enable()
	assert.Equal(t, "\033[32m\033[0m", Green(""))

	Disable()
	assert.Equal(t, "", Green(""))
	Enable()
}
