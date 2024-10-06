package cryptipass

import (
	"math"
	"math/rand/v2"
	"strings"
)

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
//   - total:   An integer representing the total count of transitions
//     for the current state, used for entropy calculations.
//   - entropy: A float64 representing the entropy of the transitions,
//     calculated using the Shannon entropy formula to provide
//     a measure of uncertainty or randomness in the transitions.
//
// The distribution struct is used internally to facilitate the generation
// of secure and unpredictable passwords by modeling the relationships
// between characters in the generated output.
type distribution struct {
	tokens  []string
	counts  []int
	total   int
	entropy float64
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
func distill(tokens []string) map[string]distribution {
	transition_matrix := make(map[string]map[rune]int)
	put := func(str string, r rune) {
		if transition_matrix[str] == nil {
			transition_matrix[str] = make(map[rune]int)
		}
		transition_matrix[str][r]++
	}
	for _, w := range tokens {
		R := []rune(strings.ToLower(w))
		if len(R) == 0 {
			continue
		}
		put("LENGTHS", rune(len(R)))
		put("", R[0])
		if len(R) == 1 {
			continue
		}
		put(string(R[0]), R[1])
		for i := 0; i < len(R)-2; i++ {
			put(string(R[i])+string(R[i+1]), R[i+2])
		}
	}
	dist_trans_matrix := make(map[string]distribution)
	for k, rfreq := range transition_matrix {
		C := 0
		for _, freq := range rfreq {
			C += freq
		}
		H := 0.0
		tr := distribution{}
		tr.counts = make([]int, 0)

		tr.tokens = make([]string, 0)
		cum := 0
		for ru, freq := range rfreq {
			p := float64(freq) / float64(C)
			H -= math.Log2(p) * p
			cum += freq
			tr.counts = append(tr.counts, cum)
			tr.tokens = append(tr.tokens, string(ru))
		}
		tr.total = C
		tr.entropy = H
		dist_trans_matrix[k] = tr
	}
	return dist_trans_matrix
}

// generator is a structure that holds the state of the password
// generator instance. It includes a random number generator (Rng)
// and a jump table (JumpTable) which defines character transition
// probabilities for generating secure, pronounceable passwords.
type generator struct {
	Rng        *rand.Rand // Rng remains public to allow custom RNGS
	jump_table *map[string]distribution
}
