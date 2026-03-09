package pkg

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

func HashPassword(text string) (string, error) {
	if text == "" {
		return "", ErrorHandler(errors.New("one of the fieds are blank"), "one of the fieds are blank")
	}

	salt := make([]byte, 16)
	_, err := rand.Read(salt)

	if err != nil {
		return "", ErrorHandler(errors.New("failed to generate salt"), "error adding data")
	}

	hash := argon2.IDKey([]byte(text), salt, 1, 64*1024, 4, 32)
	saltBase64 := base64.StdEncoding.EncodeToString(salt)
	hashBase64 := base64.StdEncoding.EncodeToString(hash)

	enodedHash := fmt.Sprintf("%s.%s", saltBase64, hashBase64)
	text = enodedHash
	return enodedHash, nil
}

func VerifyPassword(text, encodedHash string) error {
	parts := strings.Split(encodedHash, ".")
	if len(parts) != 2 {
		return ErrorHandler(errors.New("invalid encoded hash format"), "envalid encoded hash format")
	}

	saltBase64 := parts[0]
	hashedTextBase64 := parts[1]

	salt, err := base64.StdEncoding.DecodeString(saltBase64)
	if err != nil {
		return ErrorHandler(errors.New("failed to decode the salt"), "Failed to decode the salt")
	}

	hashedText, err := base64.StdEncoding.DecodeString(hashedTextBase64)
	if err != nil {
		return ErrorHandler(errors.New("failed to decode the hash password"), "Failed to decode the salt")
	}

	hash := argon2.IDKey([]byte(text), salt, 1, 64*1024, 4, 32)

	if len(hash) != len(hashedText) {
		return ErrorHandler(errors.New("incorrect length of password"), "Incorrect length of password")
	}

	if subtle.ConstantTimeCompare(hash, hashedText) == 1 {
		return nil
	}

	return ErrorHandler(errors.New("incorrect password"), "Incorrect password")

}
