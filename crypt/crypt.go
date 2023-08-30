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

// GenerateSecretKey generates a encoded string based on a 32 byte key for
// AES-256. This secret key will be used in Encrypt and Decypt functions.
// If there is an error, the error will be returned.
func GenerateSecretKey() (string, error) {
	bytes, err := generate32ByteKeyForAes256()
	if err != nil {
		return "", err
	}

	return encodeByteToString(bytes), nil
}

func generate32ByteKeyForAes256() ([]byte, error) {
	length := 32
	bytes := make([]byte, length)

	// Check the correct number of bytes
	_, err := rand.Read(bytes)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func encodeByteToString(bytes []byte) string {
	return hex.EncodeToString(bytes)
}

// Encrypt encrypts a string by the usage of a secret and the GCM
// cryptography mode. If there is an error, the error will be returned.
func Encrypt(stringToEncrypt, secret string) (string, error) {
	key, err := decodeStringToBytes(secret)
	if err != nil {
		return "", err
	}

	plainText := []byte(stringToEncrypt)

	block, err := createCipherBlock(key)
	if err != nil {
		return "", err
	}

	// Create new GCM (https://en.wikipedia.org/wiki/Galois/Counter_Mode)
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// Create a nonce from GCM
	nonce := make([]byte, aesGCM.NonceSize())

	// Check the correct number of bytes
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return "", err
	}

	// Encrypt data
	cipherText := aesGCM.Seal(nonce, nonce, plainText, nil)

	return fmt.Sprintf("%x", cipherText), nil
}

func decodeStringToBytes(data string) ([]byte, error) {
	key, err := hex.DecodeString(data)
	if err != nil {
		return nil, err
	}

	return key, nil
}

func createCipherBlock(key []byte) (cipher.Block, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	return block, nil
}

// Decrypt decrypts a encrypted string by the usage of a secret and the
// GCM cryptography mode. If there is an error, the error will be returned.
func Decrypt(encryptedString, secret string) (string, error) {
	key, err := decodeStringToBytes(secret)
	if err != nil {
		return "", err
	}

	enc, err := decodeStringToBytes(encryptedString)
	if err != nil {
		return "", err
	}

	block, err := createCipherBlock(key)
	if err != nil {
		return "", err
	}

	// Create new GCM (https://en.wikipedia.org/wiki/Galois/Counter_Mode)
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// Get the nonce size
	nonceSize := aesGCM.NonceSize()

	// Extract the nonce from the encrypted data
	nonce, cipherText := enc[:nonceSize], enc[nonceSize:]

	// Decrypt data
	plainText, err := aesGCM.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return "", err
	}

	return string(plainText), nil
}
