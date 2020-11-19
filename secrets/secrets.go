package secrets

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"

	"github.com/sharmarajdaksh/go-pwd/config"
)

// EncryptString encrypts a secret usign the master secret
func EncryptString(k string, s string) (string, error) {
	key := []byte(k)
	plaintext := []byte(s)

	ciphertext, err := encrypt(key, plaintext)
	if err != nil {
		return "", fmt.Errorf("failed to encrypt string: %w", err)
	}

	return ciphertext, nil
}

func encrypt(key, plaintext []byte) (string, error) {

	gcm, err := getGcmForKey(key)
	if err != nil {
		return "", fmt.Errorf("gcm generation failed: %w", err)
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", fmt.Errorf("failed to read nonce: %w", err)
	}

	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)
	return fmt.Sprintf("%x", ciphertext), nil
}

// DecryptString decrypts a secret using the master secret
func DecryptString(k string, s string) (string, error) {
	key := []byte(k)

	ciphertext, err := hex.DecodeString(s)
	if err != nil {
		return "", fmt.Errorf("failed to decode ciphertext: %w", err)
	}

	plaintext, err := decrypt(key, ciphertext)
	if err != nil {
		return "", fmt.Errorf("failed to decrypt string: %w", err)
	}

	return plaintext, nil

}

func decrypt(key, ciphertext []byte) (string, error) {
	gcm, err := getGcmForKey(key)
	if err != nil {
		return "", fmt.Errorf("gcm generation failed: %w", err)
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return "", fmt.Errorf("invalid ciphertext")
	}

	nonce, enc := ciphertext[:nonceSize], ciphertext[nonceSize:]

	plaintext, err := gcm.Open(nil, nonce, enc, nil)
	if err != nil {
		return "", fmt.Errorf("decrytion routine failed: %w", err)
	}

	p := string(plaintext)

	return p, nil
}

func getGcmForKey(key []byte) (cipher.AEAD, error) {
	cphr, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("failed to create aes cipher: %w", err)
	}

	gcm, err := cipher.NewGCM(cphr)
	if err != nil {
		return nil, fmt.Errorf("failed to create gcm: %w", err)
	}

	return gcm, nil
}

func getMasterSecretBytes() []byte {
	return []byte(config.GetMasterSecret())
}
