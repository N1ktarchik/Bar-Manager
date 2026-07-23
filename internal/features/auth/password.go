package auth

import (
	"crypto/sha256"
	"encoding/hex"
)

func hash(password string) string {
	hashPass := sha256.Sum256([]byte(password))

	return hex.EncodeToString(hashPass[:])
}

func (s *JWTservice) compare(password string) bool {
	hashPass := hash(password)

	return s.password == hashPass
}
