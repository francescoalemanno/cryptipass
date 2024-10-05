package main

import (
	"flag"
	"fmt"
	"math"
	"slices"
	"strings"

	cp "github.com/francescoalemanno/cryptipass"
)

type Passphrase struct {
	F string
	H float64
}

func main() {
	g := cp.NewInstance()
	f_pattern := flag.String("p", "W.w.w",
		`pattern used to generate passphrase e.g. try:
	-p WWW20dd
	
	other possible patterns are formed by combining:
	- 'W' for uppercase word.
	- 'w' lowercase word.
	- 's' symbol.
	- 'd' digit.
	- 'c' a character.
	`)
	passwords := flag.Uint64("n", 6, "number of passwords to generate")
	flag.Parse()
	pattern := *f_pattern

	pws := []Passphrase{}
	for range *passwords {
		F, H := g.GenFromPattern(pattern)
		pws = append(pws, Passphrase{F: F, H: H})
	}
	slices.SortFunc(pws, func(a, b Passphrase) int {
		if a.H > b.H {
			return 1
		} else if a.H < b.H {
			return -1
		}
		return 0
	})
	maxlen := 0
	for _, v := range pws {
		maxlen = max(maxlen, len(v.F))
	}
	title := "Passphrase"
	maxlen = max(maxlen, len(title))

	pad := 4
	padding := strings.Repeat(" ", pad+maxlen-len(title))
	fmt.Printf("%s%s%s\n\n", title, padding, "Log10(Guesses)    EntropyLog2")
	for _, p := range pws {
		padding := strings.Repeat(" ", pad+maxlen-len(p.F))
		fmt.Printf("%s%s%.2f              %.2f\n", p.F, padding, (p.H-1)/math.Log2(10), p.H)
	}
}
