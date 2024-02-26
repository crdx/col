# col

**col** is a terminal colours library for Go.

## Installation

```sh
go get crdx.org/col
```

## Usage

```go
import (
    "fmt"
    "crdx.org/col"
)

func main() {
    col.Init()
    fmt.Println(col.Green("Hello world"))
}
```

## Heuristics

If `Init` or `InitUnless` is called, the following heuristics are used to figure out whether colour support should be disabled.

- The `NO_COLOR` environment variable is set.
- The terminal is not interactive.

To disable use of these heuristics, simply don't call `Init` or `InitUnless`.

## General methods

### `Enable()`

Enable colour support (default).

### `Disable()`

Disable colour support.

### `Init()`

Initialise colour support using various heuristics.

### `InitUnless(disable bool)`

Like `Init()`, but optionally override heuristics and force colours to be disabled if `disable` is set to `true`.

For example, this may be used in the sitation where a `--no-color` command line flag has been passed and colour support should be disabled if it is true:

```go
col.InitUnless(opts.NoColor)
```

## Colour methods

### Foreground

```go
func Green(fmt any, args ...any) string
func Blue(fmt any, args ...any) string
func Red(fmt any, args ...any) string
func Yellow(fmt any, args ...any) string
func Magenta(fmt any, args ...any) string
func White(fmt any, args ...any) string
func Black(fmt any, args ...any) string
func Cyan(fmt any, args ...any) string
```

### Background

```go
func BgGreen(fmt any, args ...any) string
func BgBlue(fmt any, args ...any) string
func BgRed(fmt any, args ...any) string
func BgYellow(fmt any, args ...any) string
func BgMagenta(fmt any, args ...any) string
func BgWhite(fmt any, args ...any) string
func BgBlack(fmt any, args ...any) string
func BgCyan(fmt any, args ...any) string
```

### Styling

```go
func Underline(fmt any, args ...any) string
func Bold(fmt any, args ...any) string
func Italic(fmt any, args ...any) string
```

## Contributions

Open an [issue](https://github.com/crdx/col/issues) or send a [pull request](https://github.com/crdx/col/pulls).

## Licence

[GPLv3](LICENCE).
