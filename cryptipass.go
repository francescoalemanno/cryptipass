// Package cryptipass provides a passphrase generation and entropy calculation
// system using wordlists and patterns. It is designed to ensure randomness
// while offering flexibility in the structure of generated passwords.
package cryptipass

import (
	cr "crypto/rand"
	_ "embed"
	"fmt"
	"log"
	"math"
	"math/rand/v2"
	"strings"
)

// Transition represents the probability distribution of rune sequences within
// words. It stores the runes, their counts, the total occurrences, and the
// entropy of the sequence.
type Transition struct {
	Runes   []rune
	Counts  []int
	Total   int
	Entropy float64
}

func distill(tokens []string) map[string]Transition {
	transition_matrix := make(map[string]map[rune]int)
	put := func(str string, r rune) {
		if transition_matrix[str] == nil {
			transition_matrix[str] = make(map[rune]int)
		}
		transition_matrix[str][r]++
	}

	for _, w := range tokens {
		R := []rune(w)
		put("lengths", rune(len(R)))
		put("", R[0])
		put(string(R[0]), R[1])
		for i := 0; i < len(R)-2; i++ {
			put(string(R[i])+string(R[i+1]), R[i+2])
		}
	}
	dist_trans_matrix := make(map[string]Transition)
	for k, rfreq := range transition_matrix {
		C := 0
		for _, freq := range rfreq {
			C += freq
		}
		H := 0.0
		tr := Transition{}
		tr.Counts = make([]int, 0)

		tr.Runes = make([]rune, 0)
		cum := 0
		for ru, freq := range rfreq {
			p := float64(freq) / float64(C)
			H -= math.Log2(p) * p
			cum += freq
			tr.Counts = append(tr.Counts, cum)
			tr.Runes = append(tr.Runes, ru)
		}
		tr.Total = C
		tr.Entropy = H
		dist_trans_matrix[k] = tr
	}
	return dist_trans_matrix
}

type generator struct {
	Rng       *rand.Rand
	JumpTable *map[string]Transition
}

// NewInstanceFromList generates a new passphrase generator using a custom
// list of tokens (words). It uses the ChaCha8 random number generator seeded
// by cryptographically secure random bytes.
func NewInstanceFromList(tokens []string) *generator {
	seed := [32]byte{}
	n, err := cr.Reader.Read(seed[:])
	if err != nil || n != 32 {
		log.Fatal(n, err, seed)
	}
	rng := rand.New(rand.NewChaCha8(seed))
	g := new(generator)
	g.Rng = rng
	jtbl := distill(tokens)
	g.JumpTable = &jtbl

	return g
}

// NewInstance returns a new passphrase generator initialized with a default
// wordlist. It uses the ChaCha8 random number generator for randomization.
func NewInstance() *generator {
	return NewInstanceFromList(eff_long_word_list)
}

// GenPassphrase generates a passphrase consisting of the specified number of
// words. It returns the passphrase and its total entropy value.
func (g *generator) GenPassphrase(words uint64) (string, float64) {
	wordvec := []string{}
	total_entropy := 0.0
	for range words {
		tok, h := g.GenFromPattern("w")
		wordvec = append(wordvec, tok)
		total_entropy += h
	}
	return strings.Join(wordvec, "."), total_entropy
}

// GenFromPattern generates a passphrase following the specified pattern.
// Supported pattern symbols are:
//   - 'w', 'W': words (lowercase or capitalized)
//   - 'd': digits
//   - 's': symbols
//   - 'c', 'C': characters (lowercase or uppercase)
func (g *generator) GenFromPattern(pattern string) (string, float64) {
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
			head, h_head := g.GenWord(c)
			passphrase = passphrase + head
			entropy = entropy + h_head
		case 'd':
			d := g.Rng.IntN(10)
			H := math.Log2(10.0)
			passphrase += fmt.Sprint(d)
			entropy += H
		case 's':
			symbols := "@#!$%&=?^+-*\""
			d := g.Rng.IntN(len(symbols))
			H := math.Log2(float64(len(symbols)))
			passphrase += string(symbols[d])
			entropy += H
		case 'c', 'C':
			runc, dH := g.PickNext(strings.ToLower(passphrase))
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

// GenWord generates a word based on the character pattern (e.g., 'w', 'W').
// It returns the generated word and its entropy.
func (g *generator) GenWord(c rune) (string, float64) {
	head, h_head := g.PickNext("")
	leng, h_leng := g.PickLength()
	if c == 'W' {
		head = strings.ToUpper(head)
	}
	for len(head) < leng {
		c, h := g.PickNext(strings.ToLower(head))
		head += c
		h_head += h
	}
	h_head += h_leng
	return head, h_head
}

type CertifyResult struct {
	NominalH float64
	Gap      float64
	StdDev   float64
}

// Certify verifies the randomness and entropy of the generated passphrases.
// It runs multiple iterations to certify that the empirical entropy matches the nominal entropy.
func Certify(Gen func() (string, float64)) CertifyResult {
	nominal_H := 0.0
	nominal_H2 := 0.0
	cnt_nom_H := 0.0
	for range 1000 {

		_, nh := Gen()
		nominal_H += nh
		nominal_H2 += nh * nh
		cnt_nom_H++
	}
	cnt := make(map[string]int)
	n := float64(0)
	Q := 64
	for {
		for range Q {
			w, nh := Gen()
			nominal_H += nh
			nominal_H2 += nh * nh
			cnt_nom_H++
			cnt[w]++
			n++
		}
		Q += Q / 16
		m := float64(len(cnt))
		H := 0.0
		for _, iC := range cnt {
			c := float64(iC)
			p := (c / n)
			H -= p * math.Log2(p)
		}
		H += (m - 1) / (2 * n)
		nomH := nominal_H / cnt_nom_H
		nomH2 := nominal_H2 / cnt_nom_H
		stddev := math.Sqrt(max(nomH2-nomH*nomH, 1e-16))
		gap := nomH - H
		if math.Abs(gap) < 0.05 || math.Log2(n) > 3*nomH {
			return CertifyResult{NominalH: nomH, Gap: gap, StdDev: stddev}
		}
	}

}

// PickNext selects the next rune in the sequence based on the previous seed.
// It uses the stored transition matrix to ensure a smooth distribution of
// rune occurrences, retrying until a valid match is found.
func (g *generator) PickNext(seed string) (string, float64) {
	L := min(len(seed), 2)
	tok := strings.ToLower(seed[len(seed)-L:])
retry:
	if tr, ok := (*g.JumpTable)[tok]; ok {
		N := g.Rng.IntN(tr.Total)
		for i, v := range tr.Counts {
			if N < v {
				return string(tr.Runes[i]), tr.Entropy
			}
		}
		panic("unexpected")
	}
	tok = tok[1:]
	goto retry
}

// PickLength selects the length of the next generated word based on the
// transition matrix of word lengths. It ensures the generated word has an
// appropriate length and entropy.
func (g *generator) PickLength() (int, float64) {
	tr, ok := (*g.JumpTable)["lengths"]
	if ok {
		N := g.Rng.IntN(tr.Total)
		for i, v := range tr.Counts {
			if N < v {
				return int(tr.Runes[i]), tr.Entropy
			}
		}
	}
	panic("unexpected rand num")
}
