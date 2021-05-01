# col

A terminal colours library for Go.

## Usage

```go
import (
    "fmt"
    "github.com/crdx/col"
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

```
func Green(str string)
func Blue(str string)
func Red(str string)
func Yellow(str string)
func Magenta(str string)
func White(str string)
func Black(str string)
func Cyan(str string)
```

### Background

```
func BgGreen(str string)
func BgBlue(str string)
func BgRed(str string)
func BgYellow(str string)
func BgMagenta(str string)
func BgWhite(str string)
func BgBlack(str string)
func BgCyan(str string)
```

### Styling

```
func Underline(str string)
func Bold(str string)
func Italic(str string)
```
