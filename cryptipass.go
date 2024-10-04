// Package cryptipass v1.3.0 provides functionality for generating secure passphrases
// composed of random words, with a focus on cryptographic security and entropy calculation.
package cryptipass

import (
	cr "crypto/rand"
	"log"
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
	// Faster but non-cryptographic alternative:
	// rng = rand.New(rand.NewPCG(rng.Uint64(), rng.Uint64()))
}

// NewPassphrase generates a passphrase consisting of the specified number of random words.
// Each word is chosen by the GenMixWord function, and the total entropy of the passphrase
// is calculated and returned along with the passphrase.
//
//   - Parameters:
//     words (uint64): The number of words to include in the passphrase.
//
//   - Return values:
//     string: The generated passphrase, with words joined by periods.
//     float64: The total entropy (in bits) of the passphrase, indicating its strength.
func NewPassphrase(words uint64) (string, float64) {
	wordvec := []string{}
	total_entropy := 0.0
	for range words {
		tok, h := GenMixWord()
		wordvec = append(wordvec, tok)
		total_entropy += h
	}
	return strings.Join(wordvec, "."), total_entropy
}

// GenMixWord generates a single word of random length and returns it along with its entropy.
// The word length is determined by the PickLength function, and the word itself is generated
// by the GenWord function.
//
//   - Return values:
//     string: The generated word of random length.
//     float64: The entropy contributed by both the word length and the word itself.
func GenMixWord() (string, float64) {
	l, entropy_l := PickLength()
	w, entropy_w := GenWord(l)
	return w, entropy_l + entropy_w
}

// GenWord generates a word of exactly n characters and returns it along with the entropy
// associated with the process. The characters are selected by calling PickNext iteratively
// until the word reaches the desired length.
//
//   - Parameters:
//     n (int): The number of characters to include in the generated word.
//
//   - Return values:
//     string: The generated word consisting of n characters.
//     float64: The total entropy contributed by the character selection process.
func GenWord(n int) (string, float64) {
	s := ""
	total_entropy := 0.0
	h := 0.0
	for len(s) < n {
		s, h = PickNext(s)
		total_entropy += h
	}
	return s, total_entropy
}
