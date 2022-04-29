package goloc

import (
	"bufio"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"golang.org/x/exp/slices"
)

type File struct {
	Ext   string
	Value int
}

/*
Loads files and counts their lines.
A struct is used to reduce complexity within the filepath.WalkFunc.

Taks in a slice of strings.

	ignore := []string{}
*/
func Load(target string, ignore, extignore []string, debug bool) map[string]int {
	var sl []string
	files := []File{}
	m := map[string]int{}

	for _, s := range sl {
		ignore = append(ignore, s)
	}

	if debug {
		fmt.Printf("%s\n", strings.Repeat("-", 20))
		fmt.Printf("Total Exclusions: %d\n", len(ignore)+len(extignore))
		for i, e := range ignore {
			if e != "" {
				fmt.Printf(
					"%d. %s\n",
					i+1,
					e,
				)
			}
		}
		for i, e := range extignore {
			if e != "" {
				fmt.Printf("%d. %s\n", i+len(ignore)+1, e)
			}
		}
		fmt.Printf("%s\n", strings.Repeat("-", 20))
	}

	filepath.Walk(target, func(p string, fi fs.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
		pp := strings.Split(p, string(filepath.Separator))

		if !slices.Contains(ignore, pp[0]) {
			if !strings.HasPrefix(p, ".") && !fi.IsDir() && len(filepath.Ext(p)) != 0 {
				if !slices.Contains(extignore, filepath.Ext(p)[1:]) {
					files = append(files, File{
						Ext:   ConvExt(filepath.Ext(p)[1:]),
						Value: int(count(p)),
					})
				}

			}

		} else {
			return filepath.SkipDir
		}

		sort.Slice(files, func(i, j int) bool {
			return files[i].Value > files[j].Value
		})

		return nil
	})

	for i := range files {
		m[files[i].Ext] += files[i].Value
	}

	return m
}

// Take in io.Reader and count the number of line breaks.
func count(p string) (c int) {
	file, err := os.Open(p)
	if err != nil {
		log.Fatal(err)
	}
	sc := bufio.NewScanner(file)
	sc.Split(bufio.ScanLines)

	rgxhtml, err := regexp.Compile(`<!--[\s\S|\n].*-->`)
	if err != nil {
		log.Println(rgxhtml)
	}

	for sc.Scan() {
		if rgxhtml.MatchString(sc.Text()) {
			// do stuff
		} else {
			// don't do stuff
		}
		c++
	}

	return
}
