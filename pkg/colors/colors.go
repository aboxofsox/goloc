package colors

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"syscall"
)

const flags = 0x0001 | 0x0002 | 0x0004

type Color struct {
	FileType string `json:"file_type"`
	ASCI     string `json:"asci"`
}

var _colors = map[string]string{
	"reset":        "\033[0m",
	"javascript":   "\033[031m",  // Red
	"python":       "\033[031m",  // Red
	"typescript":   "\033[032m",  // Green
	"vue":          "\033[032m",  // Green
	"html":         "\033[33m",   // Yellow
	"css":          "\033[34m",   // Blue
	"json":         "\033[35m",   // Purple (Magenta)
	"go":           "\033[36m",   // Cyan
	"txt":          "\033[37m",   // Gray
	"bat":          "\033[97m",   // White
	"sass":         "\033[31;1m", // Bright Red
	"scss":         "\033[31;1m", // Bright Red
	"tsx":          "\033[32;1m", // Bright Green
	"ps1":          "\033[33;1m", // Bright Yellow
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
			for k := range _colors {
				delete(_colors, k)
			}

		}

	}
}

func Colorize(key, color string, n int64) string {
	return fmt.Sprintf("%s%s: %d%s", _colors[color], key, n, _colors["reset"])
}

// Generate 256 ASCI colors.
func Colors256() []string {
	var color Color
	colors := []string{}

	for i := 0; i < 16; i++ {
		for j := 0; j < 16; j++ {
			c := strconv.FormatInt(int64(i*16+j), 10)
			color.FileType = ""
			color.ASCI = fmt.Sprintf("\u001b[38;5;%sm", c)
			colors = append(colors, color.ASCI)
		}
	}

	for i := range colors {
		fmt.Printf("%s%d%s\n", colors[i], i, _colors["reset"])
	}

	return colors
}

func ColorConfig() {

}

func ljust(s, fill string, n int) string {
	return s + strings.Repeat(fill, n)
}
