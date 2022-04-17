package internal

import (
	"encoding/hex"
)
import "github.com/gtank/cryptopasta"

// Decrypt  解密
func Decrypt(text string, key string) (string, error) {
	k := &[32]byte{}
	copy(k[:], key)

	btext, _ := hex.DecodeString(text)
	ciphertext, err := cryptopasta.Decrypt(btext, k)
	if err != nil {
		return "", err
	}

	encodeText := string(ciphertext[:])
	return encodeText, nil
}
