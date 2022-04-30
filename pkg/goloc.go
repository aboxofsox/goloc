package goloc

import (
	"bufio"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/exp/slices"
)

type File struct {
	Ext   string
	Value int
}

// Load files and count their lines.
func Load(target string, ignore, extignore []string, debug bool) map[string]int {
	m := map[string]int{}

	for _, ig := range ignore {
		gs, _ := filepath.Glob(ig + "/**")
		for _, g := range gs {
			ignore = append(ignore, g)
			fmt.Println(g)
		}
	}

	filepath.Walk(target, func(p string, fi fs.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err.Error())
			return err
		}

		if slices.Contains(ignore, p) {
			return filepath.SkipDir
		}

		if !fi.IsDir() && !strings.HasPrefix(p, ".") && len(filepath.Ext(p)) != 0 && !slices.Contains(extignore, filepath.Ext(p)[1:]) {
			m[ConvExt(filepath.Ext(p)[1:])] += count(p)
		}

		return nil
	})

	return m
}

// Read file and count total number of lines.
func count(p string) (c int) {
	file, err := os.Open(p)
	if err != nil {
		log.Fatal(err)
	}
	sc := bufio.NewScanner(file)
	sc.Split(bufio.ScanLines)

	for sc.Scan() {
		c++
	}

	return
}
