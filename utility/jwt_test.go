package utility

import (
	"testing"
	"time"
)

func TestJWTParsing(t *testing.T) {
	seed := "somesecret"
	input := "someinput"
	token, err := GenerateToken(input, seed, time.Second)
	if err != nil {
		t.Fatalf("Error while generating token %s", err.Error())
	}
	id, err := ParseToken(token, seed)
	if err != nil {
		t.Fatalf("Error while parsing token %s", err.Error())
	}
	if id != input {
		t.Fatal("Token not parsed correctly")
	}
}
