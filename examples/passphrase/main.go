package main

import (
	"fmt"

	"github.com/francescoalemanno/cryptipass"
)

func main() {
	cp := cryptipass.NewInstance()

	passphrase, entropy := cp.GenPassphrase(5)
	fmt.Printf("Passphrase: %s\n", passphrase)
	fmt.Printf("Entropy: %.2f bits\n", entropy)

	passphrase, entropy = cp.GenFromPattern("W-w.cccc.CCCC(ss)[20dd]")
	fmt.Printf("Passphrase: %s\n", passphrase)
	fmt.Printf("Entropy: %.2f bits\n", entropy)

	myTokens := []string{"this", "word", "list", "is", "too", "weak", "to", "use", "in", "production", "beware", "also", "there", "is", "no", "punctuation", "since", "the", "pseudowords", "would", "be", "plagued", "by", "non", "letter", "characters"}

	gen := cryptipass.NewInstanceFromList(myTokens)
	pass, entropy := gen.GenPassphrase(4)
	fmt.Println("Custom Passphrase:", pass)
	fmt.Println("Entropy:", entropy)
}
