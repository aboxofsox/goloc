package goloc

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

func Out(str string, value int) string {
	return fmt.Sprintf("%s%s%s",
		color.CyanString(str+":"),
		strings.Repeat(".", 20-len(str)),
		color.YellowString(" %d", value),
	)
}

func OutTotal(total int) string {
	return fmt.Sprintf("Total: %s", color.GreenString("%d", total))
}

func OutNoFmt(str string, value int) string {
	return fmt.Sprintf(
		"%s%s%d",
		str,
		strings.Repeat(" ", 10-len(str)),
		value,
	)
}
