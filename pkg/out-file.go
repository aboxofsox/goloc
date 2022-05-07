package goloc

import (
	"html/template"
	"log"
	"math"
	"os"
)

var mrkdn string = `
Extension / Language               | Lines of Code                                                           |
|---------------------|-----------------------------------------------------------------------|
{{range $k, $v := .}}| **{{$k}}** | {{$v}}|
{{end}}

`

var mrkdnPBar string = `
Extension / Language               | Lines of Code                                                           |
|---------------------|-----------------------------------------------------------------------|
{{range $k, $v := .}} **{{$k}}** | ![{{$v}}](https://progress-bar.dev/{{$v}})|
{{end}}

`

// Make a markdown that includes a table.
func Mkmd(m map[string]int) {
	os.Mkdir("files", 0777)
	f, err := os.Create("gloc.md")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	tmpl, err := template.New("mrkdwn").Parse(mrkdn)
	if err != nil {
		log.Fatal(err)
	}

	err = tmpl.Execute(f, m)
	if err != nil {
		log.Fatal(err)
	}

}

// Mkade a markdown file that includes progress bars.
func MkmdPbars(m map[string]int) {
	total := 0
	prcts := map[string]float64{}
	for _, v := range m {
		total += v

	}

	for k, v := range m {
		p := math.Round((float64(v) * float64(100)) / float64(total))
		if p <= 0 {
			prcts[k] = 1
		} else {
			prcts[k] = p
		}
	}

	f, err := os.Create("gloc.md")
	if err != nil {
		log.Fatal(err)
	}

	tmpl, err := template.New("percents").Parse(mrkdnPBar)
	if err != nil {
		log.Fatal(err)
	}

	err = tmpl.Execute(f, prcts)
	if err != nil {
		log.Fatal(err)
	}
}
