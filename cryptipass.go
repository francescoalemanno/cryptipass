// Package cryptipass provides a flexible and secure password generation system that creates pronounceable passphrases based on a probabilistic model.
//
// It is designed for security-conscious developers who need to generate strong, memorable passwords.
//
// The package supports custom word lists and pattern-based password generation, allowing users to tailor the output to their needs.
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

// Transition represents a character transition model for generating
// pronounceable passwords. It encapsulates the statistical data
// necessary to understand the frequency and distribution of
// character transitions based on a given set of tokens.
//
// Fields:
//   - Runes:   A slice of runes representing characters that can follow
//     a given state in the transition matrix.
//   - Counts:  A slice of integers where each integer represents the
//     cumulative frequency count of the corresponding rune in
//     the Runes slice.
//   - Total:   An integer representing the total count of transitions
//     for the current state, used for entropy calculations.
//   - Entropy: A float64 representing the entropy of the transitions,
//     calculated using the Shannon entropy formula to provide
//     a measure of uncertainty or randomness in the transitions.
//
// The Transition struct is used internally to facilitate the generation
// of secure and unpredictable passwords by modeling the relationships
// between characters in the generated output.
type Transition struct {
	Runes   []rune
	Counts  []int
	Total   int
	Entropy float64
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
func distill(tokens []string) map[string]Transition {
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

// NewInstanceFromList creates a new instance of the cryptipass password generator
// using a custom word list. The word list should consist of tokens (words) that
// will be used to construct pronounceable passphrases.
//
// The function returns a pointer to a generator instance, which can be used
// to generate passphrases with the provided word list.
//
// Example usage:
//
//	tokens := []string{"alpha", "bravo", "charlie", "delta"}
//	gen := cryptipass.NewInstanceFromList(tokens)
//	passphrase, entropy := gen.GenPassphrase(4)
//
// Parameters:
//   - tokens []string: A list of words (tokens) from which the generator will
//     build passphrases.
//
// Returns:
//   - *generator: A new generator instance using the custom word list.
//
// If the random seed cannot be read from the crypto/rand source, the function
// will log a fatal error and terminate the application.
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

// NewInstance creates a new generator using a default word list to
// produce pronounceable, secure passphrases.
//
// The generator uses a probabilistic transition model, initialized with a pre-defined
// set of tokens to construct pronounceable passwords based on patterns of letter transitions.
//
// Example usage:
//
//	gen := cryptipass.NewInstance()
//	passphrase, entropy := gen.GenPassphrase(4)
//
// This will create a 4-word passphrase with high entropy.
//
// NewInstance is ideal for general use cases where you want to generate secure
// passphrases with a focus on ease of pronunciation and memorization.
func NewInstance() *generator {
	return NewInstanceFromList(eff_short_word_list_2_0)
}

// GenPassphrase generates a passphrase composed of a specified number of words.
//
// Each word is randomly selected based on transition probabilities, ensuring
// pronounceability while maintaining strong entropy. The words are separated by
// dots to improve readability and security (similar to the Diceware method).
//
// This method uses a cryptographically secure random number generator (via `crypto/rand`)
// for enhanced security.
//
// Parameters:
//   - words: The number of words to include in the passphrase.
//
// Returns:
//   - passphrase (string): The generated passphrase, with words separated by dots.
//   - entropy (float64): The total entropy of the passphrase, measured in bits.
//
// Example:
//
//	gen := cryptipass.NewInstance()
//	passphrase, entropy := gen.GenPassphrase(4)
//	fmt.Println("Passphrase:", passphrase)
//	fmt.Println("Entropy:", entropy)
//
// The entropy value reflects the randomness of the passphrase. The higher the entropy,
// the more secure the passphrase is, making it difficult for attackers to guess.
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

// GenFromPattern generates a password or passphrase based on a user-defined pattern.
//
// The pattern string can include special placeholders that dictate the structure
// of the generated password. Each placeholder corresponds to a specific type of character.
// The supported placeholders are:
//
//   - 'w' : Generates a word in lowercase.
//   - 'W' : Generates a word with the first letter capitalized.
//   - 'd' : Generates a random digit (0-9).
//   - 's' : Generates a random symbol from a predefined set (@#!$%&=?^+-*").
//   - 'c' : Generates a random lowercase letter.
//   - 'C' : Generates a random uppercase letter.
//   - '\\' : Escapes the next character in the pattern, allowing you to include literal values.
//
// Example usage:
//
//	g := cryptipass.NewInstance()
//	passphrase, entropy := g.GenFromPattern("Ww-d-s")
//	fmt.Println("Generated passphrase:", passphrase)
//	fmt.Println("Entropy:", entropy)
//
// This will generate a passphrase consisting of a capitalized word, a lowercase word,
// a digit, and a symbol, such as "Alpha-bravo-7-$".
//
// The function returns the generated password or passphrase and the estimated entropy
// in bits, which quantifies its strength.
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
			head, h_head := g.genword(c)
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

// genword generates a random word using the internal transition matrix and length model.
// The generated word can either be lowercase or capitalized based on the input rune.
//
// The function uses a probabilistic model to pick the initial character and
// iteratively selects subsequent characters based on prior context, ensuring
// the word is pronounceable. It also chooses the word length from a distribution
// based on real-world token data.
//
// If the input rune is 'W', the generated word will be capitalized, while
// a lowercase 'w' will generate a lowercase word.
//
// Returns:
//   - A string representing the generated word.
//   - A float64 representing the total entropy of the generated word.
//
// Example:
//
//	g := cryptipass.NewInstance()
//	word, entropy := g.genword('w')
//	fmt.Printf("Generated Word: %s, Entropy: %.2f\n", word, entropy)
func (g *generator) genword(c rune) (string, float64) {
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

// Certify evaluates the entropy and randomness of passphrases generated by a given function.
//
// The Certify function runs a comprehensive statistical analysis on the provided password generator
// function `Gen` by simulating trials of passphrase generation. It computes the average entropy
// and monitors the gap between the expected entropy and the actual entropy based on the frequency
// distribution of passphrases generated.
//
// It returns a CertifyResult struct, which includes:
// - NominalH: The nominal entropy of the passphrases, averaged over all trials.
// - Gap: The difference between the nominal entropy and the actual observed entropy.
// - StdDev: The standard deviation of the nominal entropy across trials, giving a measure of variability.
//
// This function is useful for verifying the strength and unpredictability of passphrases generated by
// custom implementations of password generators.
//
// Parameters:
// - Gen: A function that generates a passphrase and returns it alongside its entropy.
//
// Returns:
//
// - CertifyResult: A struct containing the analysis of the generator's entropy.
//
// This process continues until the gap between nominal and actual entropy is small enough,
// or a sufficient number of trials has been conducted.
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

// PickNext selects the next rune based on the current seed (context) and
// a probabilistic model derived from the transition matrix. It generates a
// rune that is most likely to follow the given string context `seed`,
// where `seed` can be up to two characters long.
//
// If the `seed` is too short or does not match any known transitions in the
// matrix, it falls back to using shorter prefixes until a match is found.
// The function retries with successively smaller parts of the `seed`
// until a suitable transition is discovered.
//
// The function returns the selected rune as a string and its entropy value.
//
// Parameters:
//   - seed: A string representing the current context (up to 2 characters).
//
// Returns:
//
//	string: The next rune, represented as a string.
//	float64: The entropy value associated with the selected rune.
//
// Example:
//
//	seed := "th"
//	nextRune, entropy := generator.PickNext(seed)
//	fmt.Printf("Next rune: %s, Entropy: %.2f\n", nextRune, entropy)
//
// Panic:
//
//	This function panics if the transition matrix is not initialized or
//	if the selection process encounters an unexpected error while choosing
//	the next rune.
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

// PickLength selects a word length based on a pre-computed probability distribution
// from the transition matrix. It uses a cryptographically secure random number generator
// to ensure unpredictable outcomes.
//
// It returns an integer representing the word length and a float64 value representing
// the entropy associated with that selection.
//
// The function panics if an unexpected random number is encountered.
//
// Example:
//
//	gen := cryptipass.NewInstance()
//	length, entropy := gen.PickLength()
//	fmt.Printf("Generated word length: %d, Entropy: %.2f\n", length, entropy)
//
// Returns:
//
//	int     - Word length selected from the transition matrix.
//	float64 - Entropy of the selected length based on its likelihood.
func (g *generator) PickLength() (int, float64) {
	tr, ok := (*g.JumpTable)["LENGTHS"]
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
