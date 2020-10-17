package utils

import (
	"crypto/sha1"
	"fmt"
)

func GenerateSHA1(params ...string) string {
	var input string

	for _, param := range params {
		input += param
	}

	hash := sha1.New()
	_, _ = hash.Write([]byte(input))
	bs := hash.Sum(nil)
	return fmt.Sprintf("%x", bs)
}
