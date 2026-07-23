package auth

import "log/slog"

type JWTservice struct {
	secretKey []byte
	password  string
	log       *slog.Logger
}

func NewJWTService(secretKey []byte, password string, log *slog.Logger) *JWTservice {
	return &JWTservice{
		secretKey: secretKey,
		password:  hash(password),
		log:       log,
	}
}
