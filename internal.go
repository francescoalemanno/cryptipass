package cryptipass

import (
	"math"
	"math/rand/v2"
	"strings"
)

// Generator is a structure that holds the state of the password
// generator instance. It includes a random number Generator (Rng)
// and a jump table (JumpTable) which defines token transition
// probabilities.
type Generator struct {
	Rng        *rand.Rand // Rng remains public to allow custom RNGS
	jump_table map[string]distribution
	depth      int
}

func (g *Generator) isready() bool {
	return g.depth > 0 && g.Rng != nil && len(g.jump_table) != 0
}
func (g *Generator) assert_ready() {
	if !g.isready() {
		panic("generator must be initialised using the NewIstance or NewCustomInstance function")
	}
}

// distribution represents a token transition model for generating
// pronounceable passwords. It encapsulates the statistical data
// necessary to understand the frequency and distribution of
// tokens transitions based on a given set of seed words.
//
// Fields:
//   - tokens:   A slice of strings representing tokens that can follow
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
// of passwords by modeling the relationships between tokens in the generated output.
type distribution struct {
	tokens    []string
	counts    []int
	entropies []float64
	total     int
}

// distill processes a list of words and builds a probabilistic transition matrix
// that models the relationships between successive tokens within those words.
//
// The transition matrix is a map where each key represents a string prefix,
// and the value is a Transition struct that tracks the possible next tokens, their frequencies,
// and the total number of transitions.
//
// This matrix enables generating new sequences of tokens based on learned patterns,
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
//	// 'transitions' will contain a probabilistic model for generating new token sequences.
func distill(tokens []string, depth int) map[string]distribution {
	transition_matrix := transition_matrix_from_tokens(tokens, depth)
	dist_trans_matrix := make(map[string]distribution)
	for k, rfreq := range transition_matrix {
		dist_trans_matrix[k] = distribution_from_counts(rfreq)
	}
	return dist_trans_matrix
}

func transition_matrix_from_tokens(tokens []string, depth int) map[string]map[string]int {
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
		for i := range R {
			from := string(R[max(i-depth, 0):i])
			to := string(R[i:min(i+depth, len(R))])
			if to == "" || from == to {
				continue
			}
			put(from, to)
		}
	}
	return transition_matrix
}

func distribution_from_counts(rfreq map[string]int) distribution {
	C := 0
	for _, freq := range rfreq {
		C += freq
	}
	tr := new_empty_distribution()
	cum := 0
	for ru, freq := range rfreq {
		p := float64(freq) / float64(C)
		cum += freq
		tr.counts = append(tr.counts, cum)
		tr.tokens = append(tr.tokens, string(ru))
		tr.entropies = append(tr.entropies, math.Log2(1.0/p))
	}
	tr.total = C
	return tr
}

func new_empty_distribution() distribution {
	tr := distribution{}
	tr.counts = make([]int, 0)
	tr.tokens = make([]string, 0)
	tr.entropies = make([]float64, 0)
	tr.total = 0
	return tr
}
