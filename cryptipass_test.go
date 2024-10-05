package cryptipass_test

import (
	"fmt"
	"math"
	"math/rand/v2"
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

func TestCert(t *testing.T) {
	g := cryptipass.Generator{}
	g.Rng = rand.New(rand.NewPCG(37512033, 27996124))

	type FN func() (string, float64)
	funcs := []FN{
		func() (string, float64) {
			return g.GenFromPattern("Cccs")
		},
		func() (string, float64) {
			return g.GenFromPattern("ddss")
		},
		func() (string, float64) {
			return g.GenFromPattern("Cdd")
		},
		func() (string, float64) {
			r := g.Rng.IntN(2)
			s, h := g.GenFromPattern(strings.Repeat("cs", 1+r))
			h += 1
			return s, h
		},
		func() (string, float64) {
			r, h := g.PickLength()
			r2, h2 := g.PickLength()
			return fmt.Sprint(r, r2), h + h2
		},
	}

	for i, x := range funcs {
		X := cryptipass.Certify(x)
		if math.Abs(X.Gap) >= 0.5 {
			t.Fatal("failed certification of function", i, X)
		} else {
			t.Log("passed certification of function", i, X, ", OK!")
		}
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
