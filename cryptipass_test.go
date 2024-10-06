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
	g := cryptipass.NewInstance()
	g.Rng = rand.New(rand.NewPCG(37512033, 27996124))

	type FN func() (string, float64)
	funcs := []FN{
		func() (string, float64) {
			return g.GenFromPattern("Ccc20d")
		},
		func() (string, float64) {
			return g.GenFromPattern("cCc2d0")
		},
		func() (string, float64) {
			return g.GenFromPattern("dds")
		},
		func() (string, float64) {
			return g.GenFromPattern("Cdd")
		},
		func() (string, float64) {
			return g.GenFromPattern("scd")
		},
		func() (string, float64) {
			N := 3
			r := g.Rng.IntN(N)
			s, h := g.GenFromPattern(strings.Repeat("c", 2+r))
			h += math.Log2(float64(N))
			return s, h
		},
		func() (string, float64) {
			r, h := g.PickLength()
			return fmt.Sprint(r), h
		},
		func() (string, float64) {
			words := []string{"", "ciao", "gi", "ok", "s"}
			N := len(words)
			r := g.Rng.IntN(N)
			nex, h := g.PickNext(words[r])
			h += math.Log2(float64(N))
			return words[r] + nex, h
		},
	}

	for i, x := range funcs {
		X := cryptipass.Certify(x)
		if math.Abs(X.Gap) >= 0.1 {
			t.Fatal("failed certification of function", i+1, X)
		} else {
			t.Log("passed certification of function", i+1)
		}
	}

}

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
