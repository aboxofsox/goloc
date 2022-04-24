package main

import (
	goloc "aboxofsox/goloc/pkg"
	"flag"
)

func main() {
	var (
		UseGitIgnore bool
		UseConfig    bool
		NoDotFiles   bool
		NoFormat     bool
		Debug        bool

		Ignore string
	)

	// Set boolean flags
	flag.BoolVar(&UseGitIgnore, "use-gitignore", false, "Choose to use .gitignore for directory exclusion.")
	flag.BoolVar(&UseConfig, "use-custom-colors", false, "Use custom colors as defined in the config.")
	flag.BoolVar(&NoDotFiles, "no-dotfiles", false, "Choose to ignore ALL dot files.")
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

		for _, g := range gi.Contents {
			tail = append(tail, g)
		}
	}

	total := 0
	fs := goloc.Load(tail, Debug)
	println()
	for ext, value := range fs {
		total += value
		if NoFormat {
			println(goloc.OutNoFmt(ext, value))
		} else {
			println(goloc.Out(ext, value))
		}
	}
	println()
}
