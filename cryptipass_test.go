package cryptipass_test

import (
	"math"
	"math/rand/v2"
	"strings"
	"testing"

	"github.com/francescoalemanno/cryptipass/v3"
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
// This process continues until the gap between nominal and actual entropy is small enough.
func Certify(Gen func() (string, float64)) CertifyResult {
	nominal_H := 0.0
	nominal_H2 := 0.0
	cnt_nom_H := 0.0
	cnt := make(map[string]int)
	n := float64(0)
	Q := 512
	ogap := 0.0
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
		H := 0.0
		for _, iC := range cnt {
			c := float64(iC)
			p := (c / n)
			H -= p * math.Log2(p)
		}
		H += (float64(len(cnt)) - 1) / (2 * n)
		nomH := nominal_H / cnt_nom_H
		nomH2 := nominal_H2 / cnt_nom_H
		stddev := math.Sqrt(max(nomH2-nomH*nomH, 1e-16))
		gap := nomH - H
		delta := ogap - gap
		ogap = gap
		if math.Abs(delta) < 0.0002 || math.Abs(gap) < 0.09 {
			return CertifyResult{NominalH: nomH, Gap: gap, StdDev: stddev}
		}
	}

}
func TestCert(t *testing.T) {
	g := cryptipass.NewCustomInstance(cryptipass.WordListDebug(), 3)
	g.Rng = rand.New(rand.NewPCG(37512033, 27996124))
	type FN func() (string, float64)
	funcs := []FN{
		func() (string, float64) {
			return g.GenFromPattern("Ccc20d")
		},
		func() (string, float64) {
			return g.GenFromPattern("ww")
		},
		func() (string, float64) {
			return g.GenFromPattern("w.w")
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
			words := []string{"", "ciao", "gi", "ok", "s"}
			N := len(words)
			r := g.Rng.IntN(N)
			nex, h := g.GenNextToken(words[r])
			h += math.Log2(float64(N))
			return words[r] + nex, h
		},
	}

	for i, x := range funcs {
		X := Certify(x)
		if math.Abs(X.Gap) >= 0.1 {
			t.Fatal("failed certification of function", i+1, X)
		} else {
			t.Log("passed certification of function", i+1, "gap:", X.Gap)
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
	next, entropy := g.GenNextToken(seed)
	if len(next) == 0 {
		t.Errorf("Expected string extension to be atleast 1 rune long")
	}

	if entropy < 0 {
		t.Errorf("Expected entropy to be >= 0, got %f", entropy)
	}
}

func TestGenFromPattern(t *testing.T) {
	g := cryptipass.NewInstance()
	pattern := "Cccc.cccc@dd"
	word, entropy := g.GenFromPattern(pattern)

	if len(word) < len(pattern) {
		t.Errorf("Expected word length > %d, got %d", len(pattern), len(word))
	}

	if entropy < 0 {
		t.Errorf("Expected entropy to be greater than 0, got %f", entropy)
	}
}

func TestPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	g := cryptipass.Generator{} //its jump-table is uninitialized
	g.GenPassphrase(2)          //this must panic
}
