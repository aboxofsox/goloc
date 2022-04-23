package colors

import (
	"fmt"
	"os"
	"runtime"
	"syscall"
)

const flags = 0x0001 | 0x0002 | 0x0004

var colors = map[string]string{
	"reset":        "\033[0m",
	"javascript":   "\033[031m", // Red
	"python":       "\033[031m", // Red
	"typescript":   "\033[032m", // Green
	"vue":          "\033[032m", // Green
	"html":         "\033[33m",  // Yellow
	"css":          "\033[34m",  // Blue
	"json":         "\033[35m",  // Purple (Magenta)
	"go":           "\033[36m",  // Cyan
	"txt":          "\033[37m",
	"white":        "\033[97m",
	"sass":         "\033[31;1m", // Bright Red
	"scss":         "\033[31;1m", // Bright Red
	"tsx":          "\033[32;1m",
	"ps1":          "\033[33;1m",
	"jsx":          "\033[34;1m", // Bright Blue
	"c":            "\033[35;1m", // Bright Purple
	"csharp":       "\033[36;1m", // Bright Cyan
	"bright_white": "\033[97;1m",
}

func Init() {
	if runtime.GOOS == "windows" {
		h := syscall.Handle(os.Stdout.Fd())
		k32dll := syscall.NewLazyDLL("kernel32.dll")
		mode := k32dll.NewProc("SetConsoleMode")

		if _, _, err := mode.Call(uintptr(h), flags); err != nil && err.Error() != "The operation completed successfully." {
			for k := range colors {
				delete(colors, k)
			}

		}

	}
}

func Color(key, color string, n int64) string {
	return fmt.Sprintf("%s%s: %d%s", colors[color], key, n, colors["reset"])
}
