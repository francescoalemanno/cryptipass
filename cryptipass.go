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
	g.jump_table = &jtbl

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
			head, h_head := g.GenNextToken("")
			leng, h_leng := g.GenWordLength()
			if c == 'W' {
				head = strings.ToUpper(head)
			}
			for len(head) < leng {
				nc, nh := g.GenNextToken(strings.ToLower(head))
				head += nc
				h_head += nh
			}
			passphrase = passphrase + head
			entropy = entropy + h_head + h_leng
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
			tok, dH := g.GenNextToken(strings.ToLower(passphrase))
			if c == 'C' {
				tok = strings.ToUpper(tok)
			}
			passphrase += string(tok)
			entropy += dH
		default:
			passphrase += string(c)
			entropy += 0.0
		}
	}
	return passphrase, entropy
}

// GenNextToken selects the next token based on the current seed (context) and
// a probabilistic model derived from the transition matrix. It generates a
// token that is most likely to follow the given string context `seed`,
// where `seed` can be up to two characters long.
//
// If the `seed` is too short or does not match any known transitions in the
// matrix, it falls back to using shorter prefixes until a match is found.
// The function retries with successively smaller parts of the `seed`
// until a suitable transition is discovered.
//
// The function returns the selected token as a string and its entropy value.
//
// Parameters:
//   - seed: A string representing the current context (up to 2 characters).
//
// Returns:
//
//	string: The next token.
//	float64: The entropy value associated with the selected token.
//
// Example:
//
//	seed := "th"
//	nextTok, entropy := generator.GenNextToken(seed)
//	fmt.Printf("Next token: %s, Entropy: %.2f\n", nextTok, entropy)
//
// Panic:
//
//	This function panics if the transition matrix is not initialized or
//	if the selection process encounters an unexpected error while choosing
//	the next token.
func (g *generator) GenNextToken(seed string) (string, float64) {
	L := min(len(seed), 2)
	tok := strings.ToLower(seed[len(seed)-L:])
retry:
	if tr, ok := (*g.jump_table)[tok]; ok {
		N := g.Rng.IntN(tr.total)
		for i, v := range tr.counts {
			if N < v {
				return tr.tokens[i], tr.entropy
			}
		}
		panic("unexpected")
	}
	tok = tok[1:]
	goto retry
}

// GenWordLength selects a word length based on a pre-computed probability distribution
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
//	length, entropy := gen.GenWordLength()
//	fmt.Printf("Generated word length: %d, Entropy: %.2f\n", length, entropy)
//
// Returns:
//
//	int     - Word length selected from the transition matrix.
//	float64 - Entropy of the selected length based on its likelihood.
func (g *generator) GenWordLength() (int, float64) {
	tr, ok := (*g.jump_table)["LENGTHS"]
	if ok {
		N := g.Rng.IntN(tr.total)
		for i, v := range tr.counts {
			if N < v {
				return int(tr.tokens[i][0]), tr.entropy
			}
		}
	}
	panic("unexpected rand num")
}
