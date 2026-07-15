package auth

import "log/slog"

type JWTservice struct {
	secretKey []byte
	log       *slog.Logger
}

func NewJWTService(secretKey []byte, log *slog.Logger) *JWTservice {
	return &JWTservice{
		secretKey: secretKey,
		log:       log,
	}
}
