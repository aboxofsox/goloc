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

// Load files into a map and call count() to count the LoC
func Load(target string, ignore, extignore []string) map[string]int {
	m := map[string]int{}

	filepath.Walk(target, func(p string, fi fs.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err.Error())
			return err
		}

		// Ignore directories listed in ignore and explicitly ignore the .git directory.
		if slices.Contains(ignore, p) && p == ".git" {
			return filepath.SkipDir
		}

		// Ignore directories, hidden directories, files with no extension, and extensions in extignore
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
	defer file.Close()

	sc := bufio.NewScanner(file)
	sc.Split(bufio.ScanLines)

	for sc.Scan() {
		c++
	}

	return
}
