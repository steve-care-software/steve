package cryptography

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/ed25519"
	crypto_rand "crypto/rand"
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/tyler-smith/go-bip39"
	"github.com/tyler-smith/go-bip39/wordlists"
	"golang.org/x/crypto/pbkdf2"
)

type application struct {
}

func createApplication() Application {
	out := application{}
	return &out
}

// Encrypt encrypts data using a password
func (app *application) Encrypt(message []byte, password []byte) ([]byte, error) {
	// Derive a key from the password using PBKDF2
	salt := make([]byte, 16)
	_, err := io.ReadFull(crypto_rand.Reader, salt)
	if err != nil {
		return nil, err
	}

	// Use PBKDF2 with SHA-256 to derive a 32-byte key from the password
	key := pbkdf2.Key([]byte(password), salt, 65536, 32, sha256.New)

	// Create a new AES cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// Use GCM for authenticated encryption
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	// Generate a random nonce for GCM
	nonce := make([]byte, aesGCM.NonceSize())
	_, err = io.ReadFull(crypto_rand.Reader, nonce)
	if err != nil {
		return nil, err
	}

	// Encrypt the plaintext using AES-GCM
	ciphertext := aesGCM.Seal(nil, nonce, message, nil)

	// Prepend the salt and nonce to the ciphertext
	fullMessage := append(salt, nonce...)
	return append(fullMessage, ciphertext...), nil
}

// Decrypt decrypts a cipher using a password
func (app *application) Decrypt(encrypted []byte, password []byte) ([]byte, error) {
	// Ensure the decoded message is long enough to contain salt, nonce, and ciphertext
	if len(encrypted) < 16+12 {
		return nil, errors.New("decoded message is too short")
	}

	// Extract the salt, nonce, and ciphertext from the decoded message
	salt := encrypted[:16]
	nonce := encrypted[16 : 16+12]
	ciphertext := encrypted[16+12:]

	// Derive the key from the password using the extracted salt
	key := pbkdf2.Key([]byte(password), salt, 65536, 32, sha256.New)

	// Create a new AES cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// Use GCM for decryption
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	// Decrypt the ciphertext using AES-GCM
	return aesGCM.Open(nil, nonce, ciphertext, nil)
}

// GeneratePrivateKey generates a private key and returns it
func (app *application) GeneratePrivateKey(language uint8) (ed25519.PrivateKey, []string, error) {
	switch language {
	case LangEnglish:
		bip39.SetWordList(wordlists.English)
	case LangFrench:
		bip39.SetWordList(wordlists.French)
	default:
		return nil, nil, errors.New("the language is invalid")
	}

	// Generate 256 bits (32 bytes) of entropy
	entropy, err := bip39.NewEntropy(256)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to generate entropy: %w", err)
	}

	// Generate the mnemonic from the entropy
	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to generate mnemonic: %w", err)
	}

	pk, err := app.generatePrivateKey(mnemonic)
	if err != nil {
		return nil, nil, err
	}

	return pk, strings.Split(mnemonic, " "), nil
}

func (app *application) GeneratePrivateKeyFromSeedWords(seedWords []string) (ed25519.PrivateKey, error) {
	// Validate mnemonic phrase
	mnemonic := strings.Join(seedWords, " ")
	return app.generatePrivateKey(mnemonic)
}

func (app *application) generatePrivateKey(mnemonic string) (ed25519.PrivateKey, error) {
	// Validate mnemonic phrase
	if !bip39.IsMnemonicValid(mnemonic) {
		return nil, errors.New("the provided seed words are invalid")
	}

	// Convert mnemonic to seed (using an optional passphrase, here empty)
	seed := bip39.NewSeed(mnemonic, "")

	// Derive the private key using PBKDF2 and SHA256 with the seed
	// Ed25519 requires a 32-byte seed, so we hash the seed with PBKDF2
	key := pbkdf2.Key(seed, []byte("ed25519 seed"), 4096, ed25519.SeedSize, sha256.New)
	return ed25519.NewKeyFromSeed(key), nil
}
