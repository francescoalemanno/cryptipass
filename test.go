package cryptipass_test

import (
	cryptipass "CryptiPass"
	"testing"
)

func TestBasic(t *testing.T) {
	pw := cryptipass.NewPassphrase(4)
	if len() < 16 {
		t.Fatalf(`Wrong length "%s"`, pw)
	}
}
