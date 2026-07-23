package middleware

import "log/slog"

type AuthService interface {
	ValidateJWT(tokenString string) error
}

type Middleware struct {
	authService AuthService
	log         *slog.Logger
}

func NewMiddleware(authService AuthService, log *slog.Logger) *Middleware {
	return &Middleware{
		authService: authService,
		log:         log,
	}
}
