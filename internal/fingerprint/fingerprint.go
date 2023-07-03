package fingerprint

import (
	"crypto/sha256"

	"github.com/google/uuid"
)

func GenerateFinderprint() (string, error) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	return uuid.String(), nil
}

func HashFingerprint(fingerprint string) (string, error) {
	hash := sha256.New()
	_, err := hash.Write([]byte(fingerprint))
	if err != nil {
		return "", err
	}
	return string(hash.Sum(nil)), nil
}
