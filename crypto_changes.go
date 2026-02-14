package go126demo

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io"
	"math/big"
)

// ============================================================
// Go 1.26: crypto/rand changes - random params now ignored
// ============================================================

// GOOD: Using crypto/rand properly
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GOOD: Using crypto/rand.Int
func GenerateRandomBigInt(max *big.Int) (*big.Int, error) {
	// In Go 1.26, the rand parameter is ignored but still accepted for compat
	return rand.Int(rand.Reader, max)
}

// BAD: Using weak hash for security context (analyzer should catch)
func WeakHash(data []byte) []byte {
	h := sha256.Sum256(data)
	return h[:]
}

// BAD: Hardcoded credentials (gosec should catch)
var apiKey = "sk-1234567890abcdef"

// BAD: Using fmt.Sprintf for error creation instead of fmt.Errorf
func BadErrorFormat(name string) error {
	return fmt.Errorf(fmt.Sprintf("user %s not found", name))
}

// GOOD: Using fmt.Errorf properly
func GoodErrorFormat(name string) error {
	return fmt.Errorf("user %s not found", name)
}

// BAD: Ignoring error from io.Copy
func BadIOCopy(dst io.Writer, src io.Reader) {
	io.Copy(dst, src)
}
