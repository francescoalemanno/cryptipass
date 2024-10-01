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

func NewPassphrase(words uint64) string {
	tokens := []string{}
	for range words {
		tokens = append(tokens, GenMixWord())
	}
	return strings.Join(tokens, ".")
}

func GenMixWord() string {
	n := rng.IntN(len(D))
	return GenWord(D[n])
}

func GenWord(n int) string {
repeat:
	s := w2[rng.IntN(len(w2))]
	for {
		next := tm[s[len(s)-2:]]
		if len(next) == 0 {
			goto repeat
		}
		c := next[rng.IntN(len(next))]
		if c == '@' {
			break
		}
		s += string(c)
		if len(s) > n {
			goto repeat
		}
	}
	if len(s) < n {
		goto repeat
	}
	return s
}
