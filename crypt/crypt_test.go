package crypt_test

import (
	"testing"

	"github.com/sven-seyfert/gomisc/crypt" //nolint:depguard
)

func TestGenerateSecretKey(t *testing.T) {
	tests := []struct {
		name              string
		expectError       bool
		expectedKeyLength int
	}{
		{name: "happy path key is generated", expectError: false, expectedKeyLength: 64},
		{name: "sad path key is not generated", expectError: true, expectedKeyLength: 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// arrange

			// act
			key, err := crypt.GenerateSecretKey()

			// assert
			if !test.expectError && err != nil {
				t.Fatalf(`unexpected error for testcase "%s"`, test.name)
			}

			if !test.expectError {
				if len(key) != test.expectedKeyLength {
					message := `generated key should have a length of %d but has a length of %d for testcase "%s"`
					t.Fatalf(message, test.expectedKeyLength, len(key), test.name)
				}
			}

			// TODO: Add sad path handling for error and key length.
		})
	}
}

func TestEncrypt(t *testing.T) {
	var listOfStringsToEncrypt = []string{
		"very secret text",
		"a longer very secret text",
		"short text",
	}

	const (
		validKey   = "38bf8b198e70e9e1d6e43c0f87ac62026b716030136e034c690546b4ba32459f" //nolint:gosec
		invalidKey = "38bf8b1"
	)

	tests := []struct {
		name                  string
		expectError           bool
		stringToEncrypt       string
		secret                string
		encryptedStringLength int
	}{
		{name: "happy path 1", expectError: false, stringToEncrypt: listOfStringsToEncrypt[0], secret: validKey, encryptedStringLength: 88},            //nolint:lll
		{name: "happy path 2", expectError: false, stringToEncrypt: listOfStringsToEncrypt[1], secret: validKey, encryptedStringLength: 106},           //nolint:lll
		{name: "happy path 3", expectError: false, stringToEncrypt: listOfStringsToEncrypt[2], secret: validKey, encryptedStringLength: 76},            //nolint:lll
		{name: "sad path invalid secret", expectError: true, stringToEncrypt: listOfStringsToEncrypt[0], secret: invalidKey, encryptedStringLength: 0}, //nolint:lll
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// arrange

			// act
			encrypted, err := crypt.Encrypt(test.stringToEncrypt, test.secret)

			// assert
			if !test.expectError && err != nil {
				t.Fatalf(`unexpected error for testcase "%s"`, test.name)
			}

			if test.expectError && err == nil {
				t.Fatalf(`a error is expected for testcase "%s"`, test.name)
			}

			if len(encrypted) != test.encryptedStringLength {
				message := `encrypted string should have a length of %d but has a length of %d for testcase "%s"`
				t.Fatalf(message, test.encryptedStringLength, len(encrypted), test.name)
			}
		})
	}
}

func TestDecrypty(t *testing.T) {
	var listOfStringsToDecrypt = []string{
		"901003d73e3efd4d950160c07dd5e1418df75394d4319825c99ecca1ce7fac892e4adaf838d4a12dff771dfc",
		"c3837efa8ec4bdb1914b05b6eae05d2c008ef144a1e3b86c34b63bad132628dbb30dfb3150b46e2596c502d5526fe280b0d9eb5225",
		"a067fc6623aaec1c2af1b01ce435377213fb50d340155fb31aa70bc856fdf58c42fc5d5fd472",
	}

	const (
		validKey   = "38bf8b198e70e9e1d6e43c0f87ac62026b716030136e034c690546b4ba32459f" //nolint:gosec
		invalidKey = "38bf8b1"
	)

	tests := []struct {
		name            string
		expectError     bool
		stringToDecrypt string
		secret          string
		decryptedString string
	}{
		{name: "happy path 1", expectError: false, stringToDecrypt: listOfStringsToDecrypt[0], secret: validKey, decryptedString: "very secret text"},          //nolint:lll
		{name: "happy path 2", expectError: false, stringToDecrypt: listOfStringsToDecrypt[1], secret: validKey, decryptedString: "a longer very secret text"}, //nolint:lll
		{name: "happy path 3", expectError: false, stringToDecrypt: listOfStringsToDecrypt[2], secret: validKey, decryptedString: "short text"},                //nolint:lll
		{name: "sad path invalid secret", expectError: true, stringToDecrypt: listOfStringsToDecrypt[0], secret: invalidKey, decryptedString: ""},              //nolint:lll
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// arrange

			// act
			decrypted, err := crypt.Decrypt(test.stringToDecrypt, test.secret)

			// assert
			if !test.expectError && err != nil {
				t.Fatalf(`unexpected error for testcase "%s"`, test.name)
			}

			if test.expectError && err == nil {
				t.Fatalf(`a error is expected for testcase "%s"`, test.name)
			}

			if decrypted != test.decryptedString {
				message := `decrypted string should be "%s" but is "%s" for testcase "%s"`
				t.Fatalf(message, test.decryptedString, decrypted, test.name)
			}
		})
	}
}
