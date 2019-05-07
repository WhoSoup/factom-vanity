package main

import (
	"crypto/rand"
	"crypto/sha256"

	"github.com/FactomProject/btcutil/base58"
	"github.com/FactomProject/ed25519"
)

const (
	keyLength   = 32 // byte size of factom keys
	checksumLen = 4  // byte size of prefixes
)

var pub = []byte{0x5f, 0xb1}  // ~FA
var priv = []byte{0x64, 0x78} // ~Fs

// Shad is the SHA-256 of the SHA-256 of the argument
func Shad(data []byte) []byte {
	h1 := sha256.Sum256(data)
	h2 := sha256.Sum256(h1[:])
	return h2[:]
}

// RCD is Factom's mechanism for addresses.
// Only supports RCD type 1
func RCD(key []byte) []byte {
	return Shad(append([]byte{0x1}, key...))
}

// HumanReadableAddress converts raw bytes to a base58 encoded string
func HumanReadableAddress(raw []byte, prefix []byte) string {
	key := append(prefix, raw[:keyLength]...)
	csum := Shad(key)[:checksumLen]
	key = append(key, csum...)
	return base58.Encode(key)
}

// RandomPair generates a random FCT private key and address combo
func RandomPair() (string, string) {
	public, private, _ := ed25519.GenerateKey(rand.Reader)
	return HumanReadableAddress(private[:], priv), HumanReadableAddress(RCD(public[:]), pub)
}
