package cryptoutils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"

	"github.com/cockroachdb/errors"
)

// AESEncrypt encrypts plaintext with AES-GCM and returns nonce||ciphertext.
// key must be a valid AES key (16, 24 or 32 bytes); pair with AESDecrypt.
func AESEncrypt(plaintext []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create cipher")
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())

	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate nonce")
	}

	ciphertext := gcm.Seal(nil, nonce, plaintext, nil)
	result := append(nonce, ciphertext...) //nolint: gocritic, makezero // as expected

	return result, nil
}

// AESDecrypt reverses AESEncrypt: it expects nonce||ciphertext produced with
// the same key and returns the plaintext, erroring on a short or tampered input.
func AESDecrypt(encrypted []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create cipher")
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create gcm")
	}

	if len(encrypted) < gcm.NonceSize() {
		return nil, errors.Errorf("encrypted data is too short, act: %d, exp: %d", len(encrypted), gcm.NonceSize())
	}

	plaintext, err := gcm.Open(nil, encrypted[:gcm.NonceSize()], encrypted[gcm.NonceSize():], nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}
