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
	cert := flag.Bool("c", false, "run entropy certification algorithm\nusers not necessarily need to concern with this detail.")
	depth := flag.Uint64("cd", 1, "certification effort.\nA larger number leads to more accurate results\nat the expense of exponentially longer completion times.")
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
	if *cert {
		certify(pattern, *depth)
	}

	pws := []Passphrase{}
	for range *passwords {
		F, H := cp.GenFromPattern(pattern)
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

func certify(pattern string, udepth uint64) {
	cnt := make(map[string]int)
	iN := 0
	Q := 128
	nominal_H := 0.0
	nominal_H2 := 0.0
	cnt_nom_H := 0.0
	for {
		Q += Q / 14
		for range Q {
			w, nh := cp.GenFromPattern(pattern)
			nominal_H += nh
			nominal_H2 += nh * nh
			cnt_nom_H++
			cnt[w]++
			iN++
		}

		n := float64(iN)
		m := float64(len(cnt))
		H := 0.0
		for _, iC := range cnt {
			c := float64(iC)
			p := (c / n)
			H -= p * math.Log2(p)
		}
		H += (m - 1) / (2 * n)
		fmt.Print("\r")
		CH := int(2 * H)
		for i := 0; i <= CH; i++ {
			fmt.Print("|")
		}
		nomH := nominal_H / cnt_nom_H
		nomH2 := nominal_H2 / cnt_nom_H
		stddev := math.Sqrt(max(nomH2-nomH*nomH, 0.0001))
		gap := nomH - H
		fmt.Printf("%v | E[H]-E=%.2f E[H] = %.2f ∂E[H] = %.2f  ", strings.Repeat(" ", 60-CH), gap, nomH, stddev)
		if gap < stddev/(1+float64(udepth)) {
			fmt.Print("\n")
			fmt.Println(H)
			break
		}
	}
}
