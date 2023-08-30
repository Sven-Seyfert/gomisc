package main

import (
	"fmt"

	"github.com/sven-seyfert/gomisc/crypt"
)

func main() {
	const stringToEncrypt = "this is a very secret text"
	fmt.Printf("StringToEncrypt: %s\n\n", stringToEncrypt)

	exampleWithGeneratedSecretKey(stringToEncrypt)
	exampleWithStaticSecretKey(stringToEncrypt)
}

func exampleWithGeneratedSecretKey(stringToEncrypt string) {
	fmt.Printf("First variant with generated secret key\n")

	// GenerateSecretKey generates a encoded string based on a 32 byte generatedKey for
	// AES-256. This secret generatedKey will be used in Encrypt and Decypt functions.
	generatedKey, err := crypt.GenerateSecretKey()
	if err != nil {
		// Do your error handling
		fmt.Println(err)
	}

	// Encrypt encrypts a string by the usage of a secret and the GCM
	// cryptography mode.
	encrypted1, err := crypt.Encrypt(stringToEncrypt, generatedKey)
	if err != nil {
		// Do your error handling
		fmt.Println(err)
	}

	fmt.Printf("  Encrypted = %s\n", encrypted1)

	// Decrypt decrypts a encrypted string by the usage of a secret and the
	// GCM cryptography mode.
	decrypted1, err := crypt.Decrypt(encrypted1, generatedKey)
	if err != nil {
		// Do your error handling
		fmt.Println(err)
	}

	fmt.Printf("  Decrypted = %s\n", decrypted1)
}

func exampleWithStaticSecretKey(stringToEncrypt string) {
	fmt.Printf("\nSecond variant with static secret key\n")

	// The following static key was extracted as return value from GenerateSecretKey
	staticKey := "38bf8b198e70e9e1d6e43c0f87ac62026b716030136e034c690546b4ba32459f"

	encrypted2, err := crypt.Encrypt(stringToEncrypt, staticKey)
	if err != nil {
		// Do your error handling
		fmt.Println(err)
	}

	fmt.Printf("  Encrypted = %s\n", encrypted2)

	decrypted2, err := crypt.Decrypt(encrypted2, staticKey)
	if err != nil {
		// Do your error handling
		fmt.Println(err)
	}

	fmt.Printf("  Decrypted = %s\n", decrypted2)
}
