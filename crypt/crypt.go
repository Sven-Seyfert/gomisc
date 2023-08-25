// Encrypt and decrypt data (strings) by the usage of a secret and GCM.
//
// Check out a usage example at https://github.com/sven-seyfert/gomisc/blob/main/examples/crypt/main.go
package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

// GenerateSecretKey generates a encoded string based on a 32 byte key for
// AES-256. This secret key will be used in Encrypt and Decypt functions.
func GenerateSecretKey() string {
	bytes := generate32ByteKeyForAes256()

	return encodeByteToString(bytes)
}

func generate32ByteKeyForAes256() []byte {
	length := 32
	bytes := make([]byte, length)

	// Check the correct number of bytes
	_, err := rand.Read(bytes)
	check(err)

	return bytes
}

func encodeByteToString(bytes []byte) string {
	return hex.EncodeToString(bytes)
}

// Encrypt encrypts a string by the usage of a secret and the GCM
// cryptography mode.
func Encrypt(stringToEncrypt, secret string) string {
	key := decodeStringToBytes(secret)
	plainText := []byte(stringToEncrypt)
	block := createCipherBlock(key)

	// Create new GCM (https://en.wikipedia.org/wiki/Galois/Counter_Mode)
	aesGCM, err := cipher.NewGCM(block)
	check(err)

	// Create a nonce from GCM
	nonce := make([]byte, aesGCM.NonceSize())

	// Check the correct number of bytes
	_, err = io.ReadFull(rand.Reader, nonce)
	check(err)

	// Encrypt data
	cipherText := aesGCM.Seal(nonce, nonce, plainText, nil)

	return fmt.Sprintf("%x", cipherText)
}

func decodeStringToBytes(data string) []byte {
	key, err := hex.DecodeString(data)
	check(err)

	return key
}

func createCipherBlock(key []byte) cipher.Block {
	block, err := aes.NewCipher(key)
	check(err)

	return block
}

// Decrypt decrypts a encrypted string by the usage of a secret and the
// GCM cryptography mode.
func Decrypt(encryptedString, secret string) string {
	key := decodeStringToBytes(secret)
	enc := decodeStringToBytes(encryptedString)
	block := createCipherBlock(key)

	// Create new GCM (https://en.wikipedia.org/wiki/Galois/Counter_Mode)
	aesGCM, err := cipher.NewGCM(block)
	check(err)

	// Get the nonce size
	nonceSize := aesGCM.NonceSize()

	// Extract the nonce from the encrypted data
	nonce, cipherText := enc[:nonceSize], enc[nonceSize:]

	// Decrypt data
	plainText, err := aesGCM.Open(nil, nonce, cipherText, nil)
	check(err)

	return string(plainText)
}
