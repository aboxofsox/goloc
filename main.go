package main

import (
	goloc "aboxofsox/goloc/pkg"
	"flag"
)

func main() {
	var (
		NoGitIgnore bool
		UseConfig   bool
		NoDotFiles  bool
		NoFormat    bool

		Exclude string
	)
	flag.BoolVar(&NoGitIgnore, "use-gitignore", false, "Choose to use .gitignore for directory exclusion.")
	flag.BoolVar(&UseConfig, "use-custom-colors", false, "Use custom colors as defined in the config.")
	flag.BoolVar(&NoDotFiles, "no-dotfiles", false, "Choose to ignore ALL dot files.")
	flag.BoolVar(&NoFormat, "no-format", false, "Print the LoC count unformatted")

	flag.StringVar(&Exclude, "exclude", "", "Add directories to exclude.")

	flag.Parse()
	var tail []string

	if Exclude != "" {
		tail = flag.Args()
		tail = append([]string{Exclude}, tail...)
	}

	// fmt.Printf("Use GitIgnore: %t\nUse Config: %t\nNot Dotfiles: %t\n", NoGitIgnore, UseConfig, NoDotFiles)

	// for i := range tail {
	// 	fmt.Println(tail[i])
	// }
	// goloc.Printer(goloc.Load())
	// colors.Colors256()

	fs := goloc.Load(tail)
	for k, v := range fs {
		println(goloc.Printer(k, v))
	}
}
