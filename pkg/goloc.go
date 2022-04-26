package goloc

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"log"
	"path/filepath"
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
func Load(target string, ignore []string, debug bool) map[string]int {
	var sl []string
	files := []File{}
	m := map[string]int{}

	for _, s := range sl {
		ignore = append(ignore, s)
	}

	if debug {
		fmt.Printf("%s\n", strings.Repeat("-", 20))
		fmt.Printf("Total Exclusions: %d\n", len(ignore))
		for i, e := range ignore {
			if e != "" {
				fmt.Printf(
					"%d. %s\n",
					i+1,
					e,
				)
			}

		}
		fmt.Printf("%s\n", strings.Repeat("-", 20))
	}

	filepath.Walk(target, func(p string, fi fs.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err.Error())
			return err
		}

		if !strings.HasPrefix(p, ".") {
			if !slices.Contains(ignore, p) {
				if !fi.IsDir() {
					files = append(files, File{
						Ext:   ConvExt(filepath.Ext(p)[1:]),
						Value: int(count(reader(p))),
					})
				}

			} else {
				return filepath.SkipDir
			}
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

// Read a file and covert it to io.Reader
func reader(p string) io.Reader {
	file, err := ioutil.ReadFile(p)
	if err != nil {
		log.Fatal(err)
	}

	reader := bytes.NewReader(file)
	return reader
}

// Take in io.Reader and count the number of line breaks.
func count(r io.Reader) (c int) {
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanLines)

	for sc.Scan() {
		c++
	}

	return c
}
