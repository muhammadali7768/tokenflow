package config

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"golang.org/x/crypto/pbkdf2"
)

// EncryptData encrypts any string data using a passphrase.
func EncryptData(data, passphrase string) (string, error) {
	salt := []byte(Config("SALT_KEY"))
	key := pbkdf2.Key([]byte(passphrase), salt, 4096, 32, sha256.New)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		return "", err
	}

	ciphertext := aesGCM.Seal(nonce, nonce, []byte(data), nil)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// DecryptData decrypts any string data using a passphrase.
func DecryptData(encryptedData, passphrase string) (string, error) {
	salt := []byte(Config("SALT_KEY"))
	key := pbkdf2.Key([]byte(passphrase), salt, 4096, 32, sha256.New)

	data, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := aesGCM.NonceSize()
	if len(data) < nonceSize {
		return "", errors.New("ciphertext too short")
	}

	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

func GenerateSessionKey() (string, error) {
	key := make([]byte, 32) // 256-bit key for AES-256
	_, err := io.ReadFull(rand.Reader, key)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(key), nil
}

func KeyToJSON(key *keystore.Key) (string, error) {
	jsonKey, err := json.Marshal(key)
	if err != nil {
		return "", err
	}
	return string(jsonKey), nil
}

func JSONToKey(jsonKey string) (*keystore.Key, error) {
	var key keystore.Key
	err := json.Unmarshal([]byte(jsonKey), &key)
	if err != nil {
		return nil, err
	}
	return &key, nil
}
