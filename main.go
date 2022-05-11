package main

import (
	"flag"
	"strings"

	goloc "github.com/aboxofsox/goloc/pkg"
)

func main() {
	var (
		UseGitIgnore bool
		IsOutFile    bool

		Repo      string
		Ignore    string
		IgnoreExt string

		fs map[string]int
	)

	// Set boolean flags
	flag.BoolVar(&UseGitIgnore, "use-gitignore", false, "Choose to use .gitignore for directory exclusion.")
	flag.BoolVar(&IsOutFile, "out-file", false, "Copy output to markdown.")
	flag.StringVar(&Repo, "repo", "", "Count the number of lines in a repo.")

	// Set string flags
	flag.StringVar(&Ignore, "ignore", "", "Add directories to be ignored.")
	flag.StringVar(&IgnoreExt, "ignore-ext", "", "Add extensions to be ignored.")

	flag.Parse()
	tail := []string{}
	exttail := []string{}

	// A goofy way of getting string values defined after the -ignore and -ignore-ext flag.
	if Ignore != "" {
		tail = flag.Args()
		tail = append([]string{Ignore}, tail...)
	}

	if IgnoreExt != "" {
		exttail = flag.Args()
		exttail = append([]string{IgnoreExt}, exttail...)
	}

	// If the -use-gitignore flag is set, use the .gitignore file as the basis of files and directories to ignore.
	if UseGitIgnore {
		gi := goloc.LoadGitIgnore(".gitignore")
		for _, g := range gi {
			if strings.HasPrefix(g, ".") {
				exttail = append(exttail, g[1:])
			}
			tail = append(tail, g)
		}
	}

	// If the -out-file flag is set, make a markdown file with LoC displayed in a table.
	if IsOutFile {
		goloc.Mkmd(fs)
	}

	// If the -repo flag is set, clone the repo in a temporary location, count the lines, and the remove the directory.
	// Else, count the LoC in the current directory.
	if len(Repo) > 0 {
		goloc.Gitter("tmp", Repo, tail)
	} else {
		fs = goloc.Load(".", tail, exttail)
		goloc.MakeTable(fs, "goloc")

	}

}
