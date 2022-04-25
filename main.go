package main

import (
	goloc "aboxofsox/goloc/pkg"
	"flag"
	"fmt"
	"sort"
	"strings"
)

func sorter(m map[string]int) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		return m[keys[i]] > m[keys[j]]
	})

	return keys
}

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

		for _, g := range gi.Contents {
			tail = append(tail, g)
		}
	}

	total := 0
	fs := goloc.Load(tail, Debug)
	fss := sorter(fs)
	println()
	for _, k := range fss {
		total += fs[k]
		if NoFormat {
			println(goloc.OutNoFmt(k, fs[k]))
		} else {
			println(goloc.Out(k, fs[k]))
		}

	}

	fmt.Println(strings.Repeat("-", 10))
	println(goloc.OutTotal(total))

	println()
}
