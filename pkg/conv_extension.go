package goloc

import (
	"strings"
)

// Convert extension name to it's associated language
func ConvExt(ext string) (res string) {
	trm := 0
	switch ext {

	case "js", "jsx", "vue", "react", "vuex":
		res = "javascript"

	case "ts", "tsx":
		res = "typescript"

	case "cs":
		res = "c#"

	case "rs":
		res = "rust"

	case "ps1":
		res = "powershell"

	case "bat":
		res = "batch"

	case "md":
		res = "markdown"

	case "yaws":
		res = "erlang"

	case "jsp", "jspx", "wss":
		res = "java"

	case "cfm":
		res = "coldfusion"

	case "asp":
		res = "asp"

	case "aspx", "axd", "asmx", "ashx":
		res = "asp.net"

	case "php", "php4", "php3", "phtml":
		res = "php"

	case "coffee", "_coffee", "cake", "cson", "iced":
		res = "coffeescript"

	case "rb":
		res = "ruby"

	case "txt":
		res = "text"

	case "py":
		res = "python"

	case "gitignore":
		res = "git"

	default:
		res = ext

	}
	trm = len(res)

	return trim(res, trm)
}

func trim(str string, n int) string {
	var ns []string
	if n >= len(str) {
		return str
	}
	sl := strings.Split(str, "")
	for i := 0; i < n; i++ {
		ns = append(ns, sl[i])
	}

	return strings.Join(ns, "")

}
