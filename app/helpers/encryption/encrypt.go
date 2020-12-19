package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"os"

	_ "github.com/joho/godotenv/autoload" //autoload env
)

// CipherKey key must be 32 chars long because block size is 16 bytes
var CipherKey = os.Getenv("SALT_KEY")

// AlgorithmNonceSize length of key
const AlgorithmNonceSize int = 12

// Encrypt encrypts plain text string into cipher text string
func Encrypt(plaintext []byte) (string, error) {
	key := []byte(CipherKey)
	// Generate a 96-bit nonce using a CSPRNG.
	nonce := make([]byte, AlgorithmNonceSize)
	_, err := rand.Read(nonce)
	if err != nil {
		return "", err
	}

	// Create the cipher and block.
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	cipher, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// Encrypt and prepend nonce.
	ciphertext := cipher.Seal(nil, nonce, plaintext, nil)
	ciphertextAndNonce := make([]byte, 0)

	ciphertextAndNonce = append(ciphertextAndNonce, nonce...)
	ciphertextAndNonce = append(ciphertextAndNonce, ciphertext...)

	return base64.StdEncoding.EncodeToString(ciphertextAndNonce), nil
}
