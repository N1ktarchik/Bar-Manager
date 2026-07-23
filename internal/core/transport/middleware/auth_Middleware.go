package middleware

import (
	"N1ktarchik/Bar-Manager/internal/core/errors"
	"N1ktarchik/Bar-Manager/internal/core/transport/response"

	"log/slog"
	"net/http"
	"strings"
)

func (m *Middleware) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			m.log.Debug("auth middleware: missing authorization header")
			response.RespondWithError(w, errors.BAD_REQUEST_ERR())
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			m.log.Debug("auth middleware: invalid authorization format", slog.String("header", authHeader))
			response.RespondWithError(w, errors.BAD_REQUEST_ERR())
			return
		}

		tokenString := parts[1]

		if err := m.authService.ValidateJWT(tokenString); err != nil {
			m.log.Debug("auth middleware: jwt validation failed", slog.Any("err", err))
			response.RespondWithError(w, err)
			return
		}

		m.log.Debug("auth middleware: user authenticated")
		next.ServeHTTP(w, r)
	})
}
