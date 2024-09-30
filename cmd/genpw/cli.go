package main

import (
	"flag"
	"fmt"
	"math"

	cp "github.com/francescoalemanno/cryptipass"
)

func main() {
	words := flag.Uint64("w", 4, "number of words per password (20.25 bits of entropy per word)")
	passwords := flag.Uint64("p", 4, "number of passwords to generate")
	flag.Parse()
	for range *passwords {
		fmt.Println(cp.NewPassphrase(*words))
	}
}

func certify() {
	cnt := make(map[string]int)
	S := 0
	for {
		cnt[cp.GenMixWord()]++
		S++
		if S%10000 == 0 {
			H := 0.0
			for _, c := range cnt {
				p := float64(c) / float64(S)
				H -= p * math.Log2(p)
			}

			fmt.Println(len(cnt), S, H)
		}
	}
}
