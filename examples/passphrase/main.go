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

}
