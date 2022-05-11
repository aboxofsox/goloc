package goloc

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/go-git/go-git/v5"
)

const BaseUrl string = "https://github.com/"

// Clone a GitHub repo into a temporary directory.
// This operation may introduce numerous read/writes, depending on the size of the cloned repo.
func Gitter(tmp, repo string, ignore []string) {
	dir, err := ioutil.TempDir("", "tmp")
	if err != nil {
		log.Fatal(err)
	}

	defer os.RemoveAll(dir)

	_, err = git.PlainClone(dir, false, &git.CloneOptions{
		URL: BaseUrl + repo,
	})

	if err != nil {
		log.Fatal(err)
	}

	fs := Load(dir, nil, nil)

	MakeTable(fs, repo)
}
