package goloc

import (
	"aboxofsox/goloc/pkg/colors"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type config struct {
	Directories []string `json:"directories"`
}

func New() {
	c := config{
		Directories: []string{},
	}

	file, _ := json.MarshalIndent(c, "", " ")
	if _, err := os.Stat("config.json"); os.IsNotExist(err) {
		ioutil.WriteFile("config.json", file, 0644)
	} else {
		return
	}
}

func loadConfig() config {
	var c config
	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(file, &c)
	return c
}

func Load() map[string]int64 {
	c := loadConfig()
	langs := map[string]int64{}

	for _, d := range c.Directories {
		err := filepath.Walk(d, func(pp string, info fs.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if info.IsDir() {
				return nil
			}

			switch filepath.Ext(info.Name())[1:] {
			case "html":
				langs["html"] += Count(Reader(pp))
			case "css":
				langs["css"] += Count(Reader(pp))
			case "js":
				langs["javascript"] += Count(Reader(pp))
			case "ts":
				langs["typescript"] += Count(Reader(pp))
			case "go":
				langs["go"] += Count(Reader(pp))
			case "vue":
				langs["vue"] += Count(Reader(pp))
			case "json":
				langs["json"] += Count(Reader(pp))
			case "jsx":
				langs["jsx"] += Count(Reader(pp))
			case "tsx":
				langs["tsx"] += Count(Reader(pp))
			case "c":
				langs["c"] += Count(Reader(pp))
			case "cs":
				langs["csharp"] += Count(Reader(pp))
			case "txt":
				langs["text"] += Count(Reader(pp))
			case "py":
				langs["python"] += Count(Reader(pp))
			default:
				langs["other"] += Count(Reader(pp))
			}

			return nil
		})

		if err != nil {
			log.Fatal(err)
		}
	}
	return langs
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

func Printer(src map[string]int64) {
	for k, v := range src {
		out := colors.Color(k, k, v)
		fmt.Println(out)
	}
}
