package main

import (
	"flag"
	"fmt"
	"math"
	"strings"

	cp "github.com/francescoalemanno/cryptipass"
)

func main() {
	cert := flag.Bool("c", false, "run entropy certification algorithm (for developers)")
	words := flag.Uint64("w", 4, "number of words per password (20.25 bits of entropy per word)")
	passwords := flag.Uint64("p", 4, "number of passwords to generate")
	flag.Parse()
	if *cert {
		certify()
	}
	for range *passwords {
		fmt.Println(cp.NewPassphrase(*words))
	}
}

func certify() {
	cnt := make(map[string]int)
	iN := 0
	Q := 128
	for {
		Q += Q / 14
		for range Q {
			cnt[cp.GenMixWord()]++
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
		gap := (m - 1) / (2 * n)
		H += gap
		fmt.Print("\r")
		CH := int(2 * H)
		for i := 0; i <= CH; i++ {
			fmt.Print("|")
		}
		fmt.Print(strings.Repeat(" ", 60-CH), H)
		if gap < 0.01 {
			fmt.Print("\n")
			fmt.Println(H)
			break
		}
	}
}
