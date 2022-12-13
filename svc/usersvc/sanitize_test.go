package usersvc

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

const (
	amogus   = "amogus"
	notapass = "notapass"
)

func TestSanitizeAndHashLessThan8(t *testing.T) {
	_, err := SanitizeAndHash(amogus)
	if err == nil {
		t.Errorf("amogus fail")
	}
}

func TestSanitizeAndHashCorrect(t *testing.T) {
	hash, err := SanitizeAndHash(notapass)
	if err != nil {
		t.Errorf("notapass fail")
	}

	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(notapass))
	if err != nil {
		t.Errorf("notapass wrong hash")
	}
}
