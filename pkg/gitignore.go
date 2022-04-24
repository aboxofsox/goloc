package goloc

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type GitIgnore struct {
	Contents []string
	Globs    []string
}

func LoadGitIgnore() *GitIgnore {
	var gitignore GitIgnore
	_, err := os.Stat(".gitignore")
	if err != nil {
		log.Fatal(err)
	}

	file, err := ioutil.ReadFile(".gitignore")
	if err != nil {
		log.Fatal(err)
	}

	content := strings.Split(string(file), "\n")
	if len(content) == 0 {
		return nil
	}

	for _, c := range content {
		if !strings.HasPrefix(c, "#") {
			gitignore.Contents = append(gitignore.Contents, c)
		}
		if strings.HasPrefix(c, "*") {
			gitignore.Globs = append(gitignore.Globs, c)
		}
	}

	return &gitignore
}
