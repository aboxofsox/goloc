package goloc

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// Read .gitignore line-by-line.
func LoadGitIgnore(p string) (ignore []string) {
	var extignore []string

	_, err := os.Stat(p)
	if err != nil {
		log.Fatal(err)
	}

	f, err := ioutil.ReadFile(p)
	if err != nil {
		log.Fatal(err)
	}

	r := bytes.NewReader(f)
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanLines)

	for sc.Scan() {
		// Add to ignore list if the line isn't a comment.
		if !strings.HasPrefix(sc.Text(), "#") {
			ignore = append(ignore, sc.Text())
		}
		// Include files by extension.
		if strings.HasPrefix(sc.Text(), ".") && sc.Text()[0:] != "/" || sc.Text()[1:] != "*" {
			extignore = append(extignore, sc.Text())
		}

	}

	return
}
