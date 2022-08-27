package shaer

import (
	"crypto/sha1"
	"fmt"
)

func StrSHA1(word string, salts ...string) string {
	bytes := []byte(word)
	for _, salt := range salts {
		bytes = BytesSHA1(bytes, salt)
	}
	return fmt.Sprintf("%x", bytes)
}

func BytesSHA1(word []byte, salts string) []byte {
	s := []byte(salts)
	p := append(word, s...)
	h := sha1.Sum(p)
	return h[:]
}
