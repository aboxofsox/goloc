package goloc

import (
	"bytes"
	"io"
	"io/fs"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

type config struct {
	Directories []string          `json:"directories"`
	Colors      map[string]string `json:"colors"`
}

func Load(exclude []string) map[string]int64 {
	var total int64 = 0
	defs := map[string]int64{}
	println(len(exclude))

	filepath.Walk(".", func(pp string, fi fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if pp[0:1] != "." && !fi.IsDir() {
			ps := strings.Split(pp, string(filepath.Separator))[0]
			if len(exclude) > 0 {
				for i := range exclude {
					if ps != exclude[i] {
						println(ps, exclude[i])
						defs[filepath.Ext(pp)[1:]] += Count(Reader(pp))
						total += Count(Reader(pp))
					} else {
						return nil
					}
				}
			}
		}

		return nil
	})

	return defs
}

func Reader(p string) io.Reader {
	file, err := ioutil.ReadFile(p)
	if err != nil {
		log.Fatal(err)
	}

	reader := bytes.NewReader(file)
	return reader
}

func Count(r io.Reader) int64 {
	b := make([]byte, 32*1024)
	var i int64 = 1
	ls := []byte{'\n'}

	for {
		c, err := r.Read(b)
		i += int64(bytes.Count(b[:c], ls))
		switch {
		case err == io.EOF:
			return i
		case err != nil:
			return i

		}
	}
}
