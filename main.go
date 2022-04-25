package main

import (
	goloc "aboxofsox/goloc/pkg"
	"flag"
)

func main() {
	var (
		UseGitIgnore bool
		NoFormat     bool
		Debug        bool

		Ignore string
	)

	// Set boolean flags
	flag.BoolVar(&UseGitIgnore, "use-gitignore", false, "Choose to use .gitignore for directory exclusion.")
	flag.BoolVar(&NoFormat, "no-format", false, "Print the LoC count unformatted")
	flag.BoolVar(&Debug, "debug", false, "Mostly used to see a visualization of the exclusion list.")

	// Set string flags
	flag.StringVar(&Ignore, "ignore", "", "Add directories to Ignore.")

	flag.Parse()
	tail := []string{}

	if Ignore != "" {
		tail = flag.Args()
		tail = append([]string{Ignore}, tail...)
	}

	if UseGitIgnore {
		gi := goloc.LoadGitIgnore()

		for _, g := range gi {
			tail = append(tail, g)
		}
	}

	fs := goloc.Load(tail, Debug)
	if NoFormat {
		goloc.OutNoFmt(fs)
	} else {
		goloc.OutBox(fs, 8)
	}
}
