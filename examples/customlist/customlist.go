package main

import (
	"fmt"

	"github.com/francescoalemanno/cryptipass/v2"
)

func main() {
	cp := cryptipass.NewCustomInstance(cryptipass.WordListLatin(), 2) //instead of 2, try 3 and 1 to see the tradeoff between fidelity to the wordlist and entropy gain.
	avgH := 0.0
	tot := 0.0
	for i := range 200000 {
		W, H := cp.GenFromPattern("W")
		avgH += H
		tot += 1
		if i < 10 {
			fmt.Println(W, H)
		}
	}
	avgH /= tot
	fmt.Println("On average this list produces ", avgH, "bits entropy per word")
}
