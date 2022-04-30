package goloc

// Convert extension name to it's associated language
func ConvExt(ext string) (res string) {
	switch ext {

	case "js":
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

	return res
}
