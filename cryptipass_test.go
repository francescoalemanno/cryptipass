package cryptipass_test

import (
	"testing"

	"github.com/francescoalemanno/cryptipass"
)

func TestBasic(t *testing.T) {
	pw := cryptipass.NewPassphrase(4)
	if len(pw) < 15 {
		t.Fatalf(`Wrong length "%s"`, pw)
	}
}
