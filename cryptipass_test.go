package cryptipass_test

import (
	"strings"
	"testing"

	"github.com/francescoalemanno/cryptipass"
)

func TestBasic(t *testing.T) {
	g := cryptipass.NewInstance()
	pw, H := g.GenPassphrase(4)
	if len(pw) < 15 {
		t.Fatalf(`Wrong length "%s"`, pw)
	}
	if H < 40 {
		// this event is so unlikely (p = 9.86588*10^-10) it must not happen.
		t.Fatalf(`Wrong entropy "%s"`, pw)
	}
}

// TestGenPassphrase tests the GenPassphrase function for generating a passphrase and validating its length.
func TestGenPassphrase(t *testing.T) {
	g := cryptipass.NewInstance()
	words := uint64(5)
	passphrase, entropy := g.GenPassphrase(words)
	wordList := strings.Split(passphrase, ".")

	if len(wordList) != int(words) {
		t.Errorf("Expected %d words, got %d", words, len(wordList))
	}

	if entropy <= 0 {
		t.Errorf("Expected entropy to be greater than 0, got %f", entropy)
	}
}

// TestGenWord tests that GenWord generates a word and returns a positive entropy value.
func TestGenWord(t *testing.T) {
	g := cryptipass.NewInstance()
	word, entropy := g.GenWord('W')

	if len(word) == 0 {
		t.Error("Expected a word, got an empty string")
	}

	if entropy < 0 {
		t.Errorf("Expected entropy to be greater than 0, got %f", entropy)
	}
}

// TestPickNext tests that PickNext generates a valid character appended to the seed and returns entropy.
func TestPickNext(t *testing.T) {
	g := cryptipass.NewInstance()
	seed := "te"
	next, entropy := g.PickNext(seed)
	if len(next) != 1 {
		t.Errorf("Expected string extension to be 1 rune long")
	}

	if entropy < 0 {
		t.Errorf("Expected entropy to be >= 0, got %f", entropy)
	}
}

// TestPickLength tests that PickLength generates a valid word length and returns entropy.
func TestPickLength(t *testing.T) {
	g := cryptipass.NewInstance()
	length, entropy := g.PickLength()

	if length < 3 || length > 9 {
		t.Errorf("Expected length to be between 3 and 9, got %d", length)
	}

	if entropy < 0 {
		t.Errorf("Expected entropy to be greater than 0, got %f", entropy)
	}
}

// TestGenFromPattern tests that GenFromPattern generates a word of a specific length and returns a positive entropy value.
func TestGenFromPattern(t *testing.T) {
	g := cryptipass.NewInstance()
	pattern := "Cccc.cccc@dd"
	word, entropy := g.GenFromPattern(pattern)

	if len(word) != len(pattern) {
		t.Errorf("Expected word length %d, got %d", len(pattern), len(word))
	}

	if entropy < 0 {
		t.Errorf("Expected entropy to be greater than 0, got %f", entropy)
	}
}
