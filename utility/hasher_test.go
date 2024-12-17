package utility

import (
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestHasher(t *testing.T) {
	input := "example"
	hash := EncryptPassword(input, bcrypt.DefaultCost)
	if !ComparePassword(input, hash) {
		t.Fatal("Hasher failed")
	}

}
