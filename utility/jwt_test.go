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

func TestJWTParsingWithWrongSeed(t *testing.T) {
	seed := "somesecret"
	input := "someinput"
	token, err := GenerateToken(input, seed, time.Second)
	if err != nil {
		t.Fatalf("Error while generating token %s", err.Error())
	}
	id, err := ParseToken(token, "wrongseed")
	if err == nil {
		t.Fatal("Token parsed with wrong seed")
	}
	if id != "" {
		t.Fatal("Token parsed with wrong seed")
	}
}

func TestExpiredToken(t *testing.T) {
	seed := "somesecret"
	input := "someinput"
	token, err := GenerateToken(input, seed, time.Second)
	if err != nil {
		t.Fatalf("Error while generating token %s", err.Error())
	}
	time.Sleep(time.Second * 2)
	id, err := ParseToken(token, seed)
	if err == nil {
		t.Fatal("Expired token parsed")
	}
	if id != "" {
		t.Fatal("Expired token parsed")
	}

}

func TestNonExpiringToken(t *testing.T) {
	seed := "somesecret"
	input := "someinput"
	token, err := GenerateNONTTLToken(input, seed)
	if err != nil {
		t.Fatalf("Error while generating token %s", err.Error())
	}
	_, err = ParseToken(token, seed)
	if err != nil {
		t.Fatalf("Error while parsing token %s", err.Error())
	}
	time.Sleep(time.Second * 2)
	id, err := ParseToken(token, seed)
	if err != nil {
		t.Fatalf("Error while parsing token %s", err.Error())
	}
	if id != input {
		t.Fatal("Token not parsed correctly")
	}
}
