package goloc

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

func Table(table map[string]int) string {
	return ""
}

func Out(str string, value int) string {
	return fmt.Sprintf("%s%s%s",
		color.CyanString(str+":"),
		strings.Repeat(".", 20-len(str)),
		color.YellowString(" %d", value),
	)
}

func OutNoFmt(str string, value int) string {
	return fmt.Sprintf(
		"%s%s%d",
		str,
		strings.Repeat(" ", 10-len(str)),
		value,
	)
}
