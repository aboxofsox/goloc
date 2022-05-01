package main

import (
	"flag"
	"strings"

	goloc "github.com/aboxofsox/goloc/pkg"
)

func main() {
	var (
		UseGitIgnore bool
		NoFormat     bool
		Debug        bool
		IsOutFile    bool

		Ignore    string
		IgnoreExt string
	)

	// Set boolean flags
	flag.BoolVar(&UseGitIgnore, "use-gitignore", false, "Choose to use .gitignore for directory exclusion.")
	flag.BoolVar(&NoFormat, "no-format", false, "Print the LoC count unformatted")
	flag.BoolVar(&IsOutFile, "out-file", false, "Copy output to markdown.")

	// Set string flags
	flag.StringVar(&Ignore, "ignore", "", "Add directories to be ignored.")
	flag.StringVar(&IgnoreExt, "ignore-ext", "", "Add extensions to be ignored.")

	flag.Parse()
	tail := []string{}
	exttail := []string{}

	if Ignore != "" {
		tail = flag.Args()
		tail = append([]string{Ignore}, tail...)
	}

	if IgnoreExt != "" {
		exttail = flag.Args()
		exttail = append([]string{IgnoreExt}, exttail...)
	}

	if UseGitIgnore {
		gi := goloc.LoadGitIgnore(".gitignore")
		for _, g := range gi {
			if strings.HasPrefix(g, ".") {
				exttail = append(exttail, g[1:])
			}
			tail = append(tail, g)
		}
	}

	fs := goloc.Load(".", tail, exttail, Debug)
	if IsOutFile {
		goloc.Mkmd(fs)
	}
	if NoFormat {
		goloc.OutNoFmt(fs)
	} else {
		goloc.MakeTable(fs)
	}

}
