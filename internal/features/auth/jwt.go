package auth

import (
	"N1ktarchik/Bar-Manager/internal/core/errors"
	"log/slog"
	"time"

	"github.com/golang-jwt/jwt"
)

func (s *JWTservice) CreateJWT(password string) (string, error) {
	s.log.Debug("creating new JWT")

	if !s.compare(password) {
		s.log.Debug("invalid password provided, JWT creation failed")
		return "", errors.INVALID_PASSWORD_ERR()
	}

	claims := jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		s.log.Error("failed to create JWT", slog.Any("err", err))
		return "", errors.INTERNAL_SERVER_ERR()
	}

	s.log.Debug("JWT created successfully")

	return tokenString, nil
}

func (s *JWTservice) ValidateJWT(tokenString string) error {
	s.log.Debug("validating JWT")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.INVALID_SIGNING_METHOD_ERR()
		}

		return []byte(s.secretKey), nil
	})

	if err != nil {
		s.log.Debug("JWT validation failed", slog.Any("err", err))
		return errors.UNAUTHORIZED_ERR()
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		s.log.Debug("JWT validated successfully")
		return nil
	}

	s.log.Debug("invalid token payload")
	return errors.UNAUTHORIZED_ERR()
}
