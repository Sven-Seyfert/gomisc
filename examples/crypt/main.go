package main

import (
	"fmt"

	"github.com/sven-seyfert/gomisc/crypt"
)

func main() {
	const stringToEncrypt = "https://github.com/sven-seyfert/gomisc/README.md"
	fmt.Printf("StringToEncrypt: %s\n\n", stringToEncrypt)
	fmt.Printf("First variant with generated secret key\n")

	// -- First variant -------------------------------------------------------
	// GenerateSecretKey generates a encoded string based on a 32 byte key for
	// AES-256. This secret key will be used in Encrypt and Decypt functions.
	key := crypt.GenerateSecretKey()

	// Encrypt encrypts a string by the usage of a secret and the GCM
	// cryptography mode.
	encrypted1 := crypt.Encrypt(stringToEncrypt, key)
	fmt.Printf("  Encrypted = %s\n", encrypted1)

	// Decrypt decrypts a encrypted string by the usage of a secret and the
	// GCM cryptography mode.
	decrypted1 := crypt.Decrypt(encrypted1, key)
	fmt.Printf("  Decrypted = %s\n", decrypted1)

	// -- Second variant ------------------------------------------------------
	fmt.Printf("\nSecond variant with static secret key\n")

	// The following static key was extracted as return value from GenerateSecretKey
	staticKey := "38bf8b198e70e9e1d6e43c0f87ac62026b716030136e034c690546b4ba32459f"

	encrypted2 := crypt.Encrypt(stringToEncrypt, staticKey)
	fmt.Printf("  Encrypted = %s\n", encrypted2)

	decrypted2 := crypt.Decrypt(encrypted2, staticKey)
	fmt.Printf("  Decrypted = %s\n", decrypted2)
}
