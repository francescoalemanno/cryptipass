package cryptipass_test

import (
	"strings"
	"testing"

	"github.com/francescoalemanno/cryptipass"
)

func TestBasic(t *testing.T) {
	pw, H := cryptipass.NewPassphrase(4)
	if len(pw) < 15 {
		t.Fatalf(`Wrong length "%s"`, pw)
	}
	if H < 40 {
		// this event is so unlikely (p = 9.86588*10^-10) it must not happen.
		t.Fatalf(`Wrong entropy "%s"`, pw)
	}
}

// TestNewPassphrase tests the NewPassphrase function for generating a passphrase and validating its length.
func TestNewPassphrase(t *testing.T) {
	words := uint64(5)
	passphrase, entropy := cryptipass.NewPassphrase(words)
	wordList := strings.Split(passphrase, ".")

	if len(wordList) != int(words) {
		t.Errorf("Expected %d words, got %d", words, len(wordList))
	}

	if entropy <= 0 {
		t.Errorf("Expected entropy to be greater than 0, got %f", entropy)
	}
}

// TestGenMixWord tests that GenMixWord generates a word and returns a positive entropy value.
func TestGenMixWord(t *testing.T) {
	word, entropy := cryptipass.GenMixWord()

	if len(word) == 0 {
		t.Error("Expected a word, got an empty string")
	}

	if entropy < 0 {
		t.Errorf("Expected entropy to be greater than 0, got %f", entropy)
	}
}

// TestGenWord tests that GenWord generates a word of a specific length and returns a positive entropy value.
func TestGenWord(t *testing.T) {
	length := 6
	word, entropy := cryptipass.GenWord(length)

	if len(word) != length {
		t.Errorf("Expected word length %d, got %d", length, len(word))
	}

	if entropy < 0 {
		t.Errorf("Expected entropy to be greater than 0, got %f", entropy)
	}
}

// TestPickNext tests that PickNext generates a valid character appended to the seed and returns entropy.
func TestPickNext(t *testing.T) {
	seed := "te"
	updatedSeed, entropy := cryptipass.PickNext(seed)

	if len(updatedSeed) <= len(seed) {
		t.Errorf("Expected updated seed to be longer, got %s", updatedSeed)
	}

	if entropy < 0 {
		t.Errorf("Expected entropy to be >= 0, got %f", entropy)
	}
}

// TestPickLength tests that PickLength generates a valid word length and returns entropy.
func TestPickLength(t *testing.T) {
	length, entropy := cryptipass.PickLength()

	if length < 3 || length > 9 {
		t.Errorf("Expected length to be between 3 and 9, got %d", length)
	}

	if entropy < 0 {
		t.Errorf("Expected entropy to be greater than 0, got %f", entropy)
	}
}
