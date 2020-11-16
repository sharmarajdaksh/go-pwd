package secrets

import (
	"testing"
)

var encryptDecryptStringTests = []string{
	"simpleststring",
	"with$pecI@lch@ractEr$",
	"with spaces",
	"with spaces $$",
	"quite_looooooooooooooooooooooooooooooooonnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnggggggg",
	"l o o o oo o o o o o o o oo  nn n n n n n n n  g g g g g g g gg ",
}

func TestEncryptDecryptString(t *testing.T) {
	key := "SUPER_SECRET_KEY_IS_SUPER_SECRET"

	for _, tc := range encryptDecryptStringTests {
		tc := tc
		t.Run(tc, func(t *testing.T) {
			t.Parallel()

			enc, err := EncryptString(key, tc)
			checkError(err)

			dec, err := DecryptString(key, enc)
			checkError(err)

			if dec != tc {
				panic("Decrypted string does not match original string")
			}
		})
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
