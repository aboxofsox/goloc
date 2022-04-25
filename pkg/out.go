package goloc

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"text/tabwriter"

	"github.com/fatih/color"
)

var (
	tlc string = "┌"
	trc string = "┐"
	hln string = "─"
	blc string = "└"
	brc string = "┘"
	vln string = "│"
)

func OutNoFmt(m map[string]int) {
	total := 0
	srt := sorter(m)

	println()
	for _, k := range srt {
		total += m[k]
		fmt.Printf("%s: %d\n", k, m[k])
	}
	fmt.Printf("%s: %d\n", "Total", total)
	println()

}

func OutBox(m map[string]int, tabsize int) {
	total := 0
	mx := max(m)
	srt := sorter(m)
	mxl := maxlength(srt)
	w := tabwriter.NewWriter(os.Stdout, 1, tabsize, 1, ' ', 0)

	println()
	fmt.Fprintf(
		w,
		"%v\t%s\t%s%v%v\n",
		tlc,
		color.CyanString("Ext"),
		color.CyanString("LoC"),
		strings.Repeat(" ", gap(mx, 2)),
		trc,
	)

	for _, k := range srt {
		total += m[k]
		fmt.Fprintf(
			w,
			"%v\t%s:\t%d\t%v\n",
			vln,
			color.YellowString(k), m[k],
			vln,
		)
	}

	fmt.Fprintf(
		w,
		"%v%v%v%v\n",
		blc,
		strings.Repeat(hln, mxl+2),
		strings.Repeat(hln, len(strconv.FormatInt(int64(mx), 10))+2),
		brc,
	)

	w.Flush()

	fmt.Printf("%s: %d", color.HiBlueString("Total"), total)
	println()
}

func sorter(m map[string]int) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		return m[keys[i]] > m[keys[j]]
	})

	return keys
}

func max(mp map[string]int) int {
	var vs []int
	for _, v := range mp {
		vs = append(vs, v)
	}
	m := 0
	if len(vs) == 0 {
		m = 0
	}
	for i, e := range vs {
		if i == 0 || e > m {
			m = e
		}
	}

	return m
}

func maxlength(sl []string) int {
	m := 0
	if len(sl) == 0 {
		m = 0
	}
	for i, s := range sl {
		if i == 0 || len(s) > m {
			m = len(s)
		}
	}
	return m
}

func gap(mx, diff int) int {
	if diff > mx {
		diff = mx
	}
	return len(strconv.FormatInt(int64(mx), 10)) - diff
}
