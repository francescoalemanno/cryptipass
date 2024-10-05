// Package cryptipass v1.3.0 provides functionality for generating secure passphrases
// composed of random words, with a focus on cryptographic security and entropy calculation.
package cryptipass

import (
	cr "crypto/rand"
	"fmt"
	"log"
	"math"
	"math/rand/v2"
	"strings"
)

var rng *rand.Rand

// init initializes the cryptographically secure random number generator (RNG).
// It reads 32 bytes of entropy from crypto/rand and uses them to seed a ChaCha8-based RNG.
// If the required number of bytes is not read, it logs a fatal error.
func init() {
	seed := [32]byte{}
	n, err := cr.Reader.Read(seed[:])
	if err != nil || n != 32 {
		log.Fatal(n, err, seed)
	}
	rng = rand.New(rand.NewChaCha8(seed))
}

// NewPassphrase generates a passphrase consisting of the specified number of random words.
// Each word is chosen by the GenMixWord function, and the total entropy of the passphrase
// is calculated and returned along with the passphrase.
//
//   - Parameters:
//
//     words (uint64): The number of words to include in the passphrase.
//
//   - Return values:
//
//     string: The generated passphrase, with words joined by periods.
//
//     float64: The total entropy (in bits) of the passphrase, indicating its strength.
func GenPassphrase(words uint64) (string, float64) {
	wordvec := []string{}
	total_entropy := 0.0
	for range words {
		tok, h := GenFromPattern("w")
		wordvec = append(wordvec, tok)
		total_entropy += h
	}
	return strings.Join(wordvec, "."), total_entropy
}

// GenFromPattern generates a passphrase based on a specified pattern.
// The pattern dictates the structure of the generated passphrase, with the following
// format options:
//
// - `w`: Lowercase word generated via GenMixWord.
//
// - `W`: Capitalized word generated via GenMixWord (first letter uppercased).
//
// - `d`: Random digit (0-9).
//
// - `s`: Random symbol from a predefined set (@#!$%&=?^+-*").
//
// Any other characters in the pattern are appended as-is.
//
//   - Parameters:
//
//     pattern (string): A string pattern that specifies the structure of the passphrase.
//
//   - Return values:
//
//     string: The generated passphrase based on the given pattern.
//
//     float64: The total entropy (in bits) of the generated passphrase.

func GenFromPattern(pattern string) (string, float64) {
	passphrase := ""
	entropy := 0.0
	pushnext := false
	for _, c := range pattern {
		if pushnext {
			pushnext = false
			passphrase += string(c)
			continue
		}
		switch c {
		case '\\':
			pushnext = true
			continue
		case 'w', 'W':
			head, h_head := GenWord(c)
			passphrase = passphrase + head
			entropy = entropy + h_head
		case 'd':
			d := rng.IntN(10)
			H := math.Log2(10.0)
			passphrase += fmt.Sprint(d)
			entropy += H
		case 's':
			symbols := "@#!$%&=?^+-*\""
			d := rng.IntN(len(symbols))
			H := math.Log2(float64(len(symbols)))
			passphrase += string(symbols[d])
			entropy += H
		case 'c', 'C':
			runc, dH := PickNext(strings.ToLower(passphrase))
			if c == 'C' {
				runc = strings.ToUpper(runc)
			}
			passphrase += string(runc)
			entropy += dH
		default:
			passphrase += string(c)
			entropy += 0.0
		}
	}
	return passphrase, entropy
}

func GenWord(c rune) (string, float64) {
	head, h_head := PickNext("")
	leng, h_leng := PickLength()
	if c == 'W' {
		head = strings.ToUpper(head)
	}
	for len(head) < leng {
		c, h := PickNext(strings.ToLower(head))
		head += c
		h_head += h
	}
	h_head += h_leng
	return head, h_head
}
