package cryptipass

import (
	"math"
	"math/rand/v2"
	"strings"
)

// Generator is a structure that holds the state of the password
// generator instance. It includes a random number Generator (Rng)
// and a jump table (JumpTable) which defines character transition
// probabilities.
type Generator struct {
	Rng        *rand.Rand // Rng remains public to allow custom RNGS
	jump_table map[string]distribution
	depth      int
}

func (g *Generator) isready() bool {
	return g.Rng != nil && len(g.jump_table) != 0
}
func (g *Generator) assert_ready() {
	if !g.isready() {
		panic("generator must be initialised using the NewIstance or NewCustomInstance function")
	}
}

// distribution represents a character transition model for generating
// pronounceable passwords. It encapsulates the statistical data
// necessary to understand the frequency and distribution of
// character transitions based on a given set of tokens.
//
// Fields:
//   - tokens:   A slice of strings representing characters that can follow
//     a given state in the transition matrix.
//   - counts:  A slice of integers where each integer represents the
//     cumulative frequency count of the corresponding token in
//     the tokens slice.
//   - entropies: A slice of float64 representing the entropy of the transitions,
//     calculated using the Shannon entropy formula to provide
//     a measure of uncertainty or randomness in the transitions.
//   - total:   An integer representing the total count of transitions
//     for the current state, used for entropy calculations.
//
// The distribution struct is used internally to facilitate the generation
// of passwords by modeling the relationships between characters in the generated output.
type distribution struct {
	tokens    []string
	counts    []int
	entropies []float64
	total     int
}

// distill processes a list of words and builds a probabilistic transition matrix
// that models the relationships between successive characters within those words.
//
// The transition matrix is a map where each key represents a string prefix (up to two characters),
// and the value is a Transition struct that tracks the possible next characters, their frequencies,
// and the total number of transitions.
//
// This matrix enables generating new sequences of characters based on learned patterns,
// making it useful for generating pronounceable words.
//
// The function returns a map of transitions that can be used for creating passphrases or similar text generation.
//
// Parameters:
//
//	tokens - a slice of strings representing the input tokens from which to distill the transition matrix.
//	depth - the correlation length used to generate the transition matrix.
//
// Returns:
//
//	map[string]Transition - a map where each key is a string prefix and each value is a Transition struct.
//
// Example:
//
//	tokens := []string{"hello", "world", "golang"}
//	transitions := distill(tokens)
//
//	// 'transitions' will contain a probabilistic model for generating new character sequences.
func distill(tokens []string, depth int) map[string]distribution {
	transition_matrix := make(map[string]map[string]int)
	put := func(str string, r string) {
		if transition_matrix[str] == nil {
			transition_matrix[str] = make(map[string]int)
		}
		transition_matrix[str][r]++
	}
	for _, w := range tokens {
		R := []rune(strings.ToLower(w))
		if len(R) == 0 {
			continue
		}
		put("LENGTHS", string(rune(len(R))))
		for i := 0; i < min(depth, len(R)); i++ {
			put(string(R[:i]), string(R[i]))
		}
		for i := 0; i < len(R)-depth; i++ {
			put(string(R[i:i+depth]), string(R[i+depth]))
		}
	}
	dist_trans_matrix := make(map[string]distribution)
	for k, rfreq := range transition_matrix {
		C := 0
		for _, freq := range rfreq {
			C += freq
		}
		tr := distribution{}
		tr.counts = make([]int, 0)

		tr.tokens = make([]string, 0)
		cum := 0
		for ru, freq := range rfreq {
			p := float64(freq) / float64(C)
			cum += freq
			tr.counts = append(tr.counts, cum)
			tr.tokens = append(tr.tokens, string(ru))
			tr.entropies = append(tr.entropies, math.Log2(1.0/p))
		}
		tr.total = C
		dist_trans_matrix[k] = tr
	}
	return dist_trans_matrix
}
