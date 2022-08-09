package go_fireblocks

import (
	"fmt"
	"testing"
)

func TestSha256hash(t *testing.T) {
	hash, err := Sha256([]byte("helloworld"))
	if err != nil {
		panic(err)
	}
	fmt.Println(hash)
}
