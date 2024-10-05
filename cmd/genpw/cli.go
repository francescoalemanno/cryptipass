package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
	"text/tabwriter"

	"github.com/francescoalemanno/cryptipass"
)

type Passphrase struct {
	F string
	H float64
}

func main() {
	g := cryptipass.NewInstance()
	f_pattern := flag.String("p", "W.w.w",
		`pattern used to generate passphrase e.g. try:
	-p WWW20dd
	
	other possible patterns are formed by combining:
	- 'w' lowercase word, 'W' for uppercase word.
	- 'c' a lowercase character, 'C' a uppercase character.
	- 's' symbol, 'd' digit.
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
		if a.F > b.F {
			return 1
		} else if a.F < b.F {
			return -1
		}
		return 0
	})

	w := tabwriter.NewWriter(os.Stdout, 1, 1, 4, ' ', 0)
	fmt.Fprintln(w, "Passphrase\tLog10(Guesses)\tLog2Entropy\t  Strength")
	fmt.Fprintln(w, " \t \t \t ")
	for _, p := range pws {
		log10guess := (p.H - 1) / math.Log2(10)
		meterbar := int(min(log10guess/2+0.5, 12))
		meter := "[" + strings.Repeat("=", meterbar) + strings.Repeat(".", 12-meterbar) + "]"
		fmt.Fprintf(w, "%v\t   %.2f\t   %.2f\t%v\n", p.F, log10guess, p.H, meter)
	}
	w.Flush()

}
