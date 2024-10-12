package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/francescoalemanno/cryptipass/v3"
)

type Passphrase struct {
	F string
	H float64
}

func main() {
	f_pattern := flag.String("p", "W.w.w",
		`pattern used to generate passphrase e.g. try:
	-p WWW20dd
	
	other possible patterns are formed by combining:
	- 'w' lowercase word, 'W' for uppercase word.
	- 'c' a lowercase token, 'C' a uppercase token.
	- 's' symbol, 'd' digit.
	`)
	passwords := flag.Uint64("n", 6, "number of passwords to generate")
	depth := flag.Uint64("d", 2, "markov chain depth, higher values lead to more plausible words but lower entropy")
	style := flag.String("s", "eff", "pseudo-word style: eff (english), latin, italian")
	flag.Parse()
	var wlist []string

	switch *style {
	case "eff", "english":
		wlist = cryptipass.WordListEFF()
	case "latin":
		wlist = cryptipass.WordListLatin()
	case "italian":
		wlist = cryptipass.WordListItalian()
	default:
		log.Fatal("word style `", *style, "` is unknown.")
	}
	g := cryptipass.NewCustomInstance(wlist, int(*depth))

	pattern := *f_pattern

	w := tabwriter.NewWriter(os.Stdout, 1, 1, 4, ' ', 0)
	fmt.Fprintln(w, "Passphrase\tLog10(Guesses)\tLog2Entropy\t  Strength")
	fmt.Fprintln(w, " \t \t \t ")
	N := *passwords
	for N > 0 {
		F, H := g.GenFromPattern(pattern)
		N--
		log10guess := (H - 1) / math.Log2(10)
		meterbar := int(min(log10guess/2+0.5, 12))
		meter := "[" + strings.Repeat("=", meterbar) + strings.Repeat(".", 12-meterbar) + "]"
		fmt.Fprintf(w, "%v\t   %.2f\t   %.2f\t%v\n", F, log10guess, H, meter)
	}
	w.Flush()

}
