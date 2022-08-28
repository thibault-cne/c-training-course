package questions

import (
	"fmt"
	"strings"
)

const (
	S = "\033[0m"

	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	White  = "\033[37m"
)

func display(text string, space int) {
	spaces := strings.Repeat("	", space)

	fmt.Println(spaces, text)
}

func colorize(text string, color string) string {
	return string(color) + text + S
}
