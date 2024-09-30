package cryptipass_test

import (
	"cryptipass"
	"testing"
)

func TestBasic(t *testing.T) {
	pw := cryptipass.NewPassphrase(4)
	if len(pw) < 16 {
		t.Fatalf(`Wrong length "%s"`, pw)
	}
}
