package go_fireblocks

import (
	"crypto/sha256"
	"fmt"
)

func Sha256(data []byte) (string, error) {
	sum := sha256.Sum256(data)
	return fmt.Sprintf("%x", sum), nil
}
