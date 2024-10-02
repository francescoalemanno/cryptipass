package cryptipass

import (
	cr "crypto/rand"
	"log"
	"math/rand/v2"
	"strings"
)

var rng *rand.Rand

func init() {
	seed := [32]byte{}
	n, err := cr.Reader.Read(seed[:])
	if err != nil || n != 32 {
		log.Fatal(n, err, seed)
	}
	rng = rand.New(rand.NewChaCha8(seed))
	// Faster but non-cryptographic
	// rng = rand.New(rand.NewPCG(rng.Uint64(), rng.Uint64()))
}

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

func GenMixWord() (string, float64) {
	l, entropy_l := PickLength()
	w, entropy_w := GenWord(l)
	return w, entropy_l + entropy_w
}

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
