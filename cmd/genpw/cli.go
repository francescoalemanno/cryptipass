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
	cert := flag.Bool("c", false, "run entropy certification algorithm (for developers)")
	words := flag.Uint64("w", 4, "number of words per password")
	passwords := flag.Uint64("p", 4, "number of passwords to generate")
	flag.Parse()
	if *cert {
		certify()
	}

	pws := []Passphrase{}
	for range *passwords {
		F, H := cp.NewPassphrase(*words)
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

	fmt.Printf("       ENTROPY |          PASSPHRASE\n")
	fmt.Printf("+++++++++++++++|++++++++++++++++++++++++++++++++\n")
	for _, p := range pws {
		fmt.Printf("% 13.2f  |  %s\n", p.H, p.F)
	}
}

func certify() {
	cnt := make(map[string]int)
	iN := 0
	Q := 128
	nominal_H := 0.0
	cnt_nom_H := 0.0
	for {
		Q += Q / 14
		for range Q {
			w, nh := cp.GenMixWord()
			nominal_H += nh
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
		gap := nomH - H
		fmt.Printf("%v | E[H]-E=%.2f E[H] = %.2f  ", strings.Repeat(" ", 60-CH), gap, nomH)
		if gap < 0.01 {
			fmt.Print("\n")
			fmt.Println(H)
			break
		}
	}
}
