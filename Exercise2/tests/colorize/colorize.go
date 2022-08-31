package colorize

import (
	"fmt"
	"strconv"
	"strings"
)

const escape = "\033"

// Base attributes
const (
	Reset int = iota
	Bold
	Faint
	Italic
	Underline
	BlinkSlow
	BlinkRapid
	ReverseVideo
	Concealed
	CrossedOut
)

// Foreground text colors
const (
	FgBlack int = iota + 30
	FgRed
	FgGreen
	FgYellow
	FgBlue
	FgMagenta
	FgCyan
	FgWhite
)

// Foreground Hi-Intensity text colors
// Only available for terminals that support aixterm specification
const (
	FgHiBlack int = iota + 90
	FgHiRed
	FgHiGreen
	FgHiYellow
	FgHiBlue
	FgHiMagenta
	FgHiCyan
	FgHiWhite
)

// Background text colors
const (
	BgBlack int = iota + 40
	BgRed
	BgGreen
	BgYellow
	BgBlue
	BgMagenta
	BgCyan
	BgWhite
)

// Background Hi-Intensity text colors
// Only available for terminals that support aixterm specification
const (
	BgHiBlack int = iota + 100
	BgHiRed
	BgHiGreen
	BgHiYellow
	BgHiBlue
	BgHiMagenta
	BgHiCyan
	BgHiWhite
)

type Color struct {
	params []int
}

func New(value ...int) *Color {
	c := &Color{
		params: make([]int, 0),
	}

	c.params = append(c.params, value...)

	return c
}

func (c *Color) sequence() string {
	format := make([]string, len(c.params))
	for i, v := range c.params {
		format[i] = strconv.Itoa(int(v))
	}

	return strings.Join(format, ";")
}

func (c *Color) format() string {
	return fmt.Sprintf("%s[%sm", escape, c.sequence())
}

func (c *Color) unformat() string {
	return fmt.Sprintf("%s[%dm", escape, Reset)
}

func (c *Color) Set() {
	fmt.Printf("%s", c.format())
}

func (c *Color) Unset() {
	fmt.Printf("%s", c.unformat())
}

func (c *Color) Print(a ...interface{}) {
	c.Set()
	defer c.Unset()

	fmt.Print(a...)
}

func (c *Color) Printf(format string, a ...interface{}) {
	c.Set()
	defer c.Unset()

	fmt.Printf(format, a...)
}

func (c *Color) Sprintf(format string, a ...interface{}) string {
	return fmt.Sprintf(fmt.Sprintf("%s%s%s", c.format(), format, c.unformat()), a...)
}

// Shortcuts to print directly formated text to a specified color.

func Black(format string, a ...interface{}) {
	c := New(FgBlack)

	c.Printf(format, a...)
}

func Red(format string, a ...interface{}) {
	c := New(FgRed)

	c.Printf(format, a...)
}

func Green(format string, a ...interface{}) {
	c := New(FgGreen)

	c.Printf(format, a...)
}

func Yellow(format string, a ...interface{}) {
	c := New(FgYellow)

	c.Printf(format, a...)
}

func Blue(format string, a ...interface{}) {
	c := New(FgBlue)

	c.Printf(format, a...)
}

func Magenta(format string, a ...interface{}) {
	c := New(FgMagenta)

	c.Printf(format, a...)
}

func Cyan(format string, a ...interface{}) {
	c := New(FgCyan)

	c.Printf(format, a...)
}

func White(format string, a ...interface{}) {
	c := New(FgWhite)

	c.Printf(format, a...)
}

// Shortcuts to print directly formated text to a specified color.

func BlackString(format string, a ...interface{}) string {
	c := New(FgBlack)

	return c.Sprintf(format, a...)
}

func RedString(format string, a ...interface{}) string {
	c := New(FgRed)

	return c.Sprintf(format, a...)
}

func GreenString(format string, a ...interface{}) string {
	c := New(FgGreen)

	return c.Sprintf(format, a...)
}

func YellowString(format string, a ...interface{}) string {
	c := New(FgYellow)

	return c.Sprintf(format, a...)
}

func BlueString(format string, a ...interface{}) string {
	c := New(FgBlue)

	return c.Sprintf(format, a...)
}

func MagentaString(format string, a ...interface{}) string {
	c := New(FgMagenta)

	return c.Sprintf(format, a...)
}

func CyanString(format string, a ...interface{}) string {
	c := New(FgCyan)

	return c.Sprintf(format, a...)
}

func WhiteString(format string, a ...interface{}) string {
	c := New(FgWhite)

	return c.Sprintf(format, a...)
}
