package goloc

import (
	"fmt"
	"strings"
)

func ljust(s, fill string, n int) string {
	return s + strings.Repeat(fill, n)
}

func Printer(ext string, count int64) string {
	return fmt.Sprintf("%s%d", ljust(ext+":", " ", 1), count)
}
