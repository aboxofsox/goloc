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
		fmt.Printf("%s %d\n", k, m[k])
	}
	fmt.Printf("%s: %d\n", "Total", total)
	println()

}

func OutBox(m map[string]int, tabsize int) {
	total := 0
	mx := len(strconv.FormatInt(int64(max(m)), 10))
	srt := sorter(m)
	spc := 0
	w := tabwriter.NewWriter(os.Stdout, 0, tabsize, 0, ' ', 0)

	fmt.Fprintf(
		w,
		"%v\t%s%s%s%v\n",
		tlc,
		color.CyanString("Ext"),
		strings.Repeat(hln, tabsize+mx-2),
		color.CyanString("LoC"),
		trc,
	)

	for _, k := range srt {
		if len(k) < 10 {
			spc = 2
		} else if len(k) >= 10 {
			spc = 0
		}
		total += m[k]
		fmt.Fprintf(
			w,
			"%v\t%s:%v \t%d\t%v\t\n",
			vln,
			color.HiYellowString(ConvExt(k)),
			strings.Repeat(" ", spc),
			m[k],
			vln,
		)
	}

	fmt.Fprintf(
		w,
		"%v\t%s%s%d%v\n",
		blc,
		color.HiBlueString("Total"),
		strings.Repeat(hln, (tabsize+mx)-lennum(total)-1),
		total,
		brc,
	)

	w.Flush()

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

func gap(mx int) int {
	return len(strconv.FormatInt(int64(mx), 10))
}

func lennum(n int) int {
	return len(strconv.FormatInt(int64(n), 10))
}
