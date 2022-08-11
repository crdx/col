# col

**col** is a terminal colours library for Go.

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

```go
func Green(str string) string
func Blue(str string) string
func Red(str string) string
func Yellow(str string) string
func Magenta(str string) string
func White(str string) string
func Black(str string) string
func Cyan(str string) string
```

### Background

```go
func BgGreen(str string) string
func BgBlue(str string) string
func BgRed(str string) string
func BgYellow(str string) string
func BgMagenta(str string) string
func BgWhite(str string) string
func BgBlack(str string) string
func BgCyan(str string) string
```

### Styling

```go
func Underline(str string) string
func Bold(str string) string
func Italic(str string) string
```

## Contributions

Open an [issue](https://github.com/crdx/col/issues) or send a [pull request](https://github.com/crdx/col/pulls).

## Licence

[GPLv3](LICENCE).
