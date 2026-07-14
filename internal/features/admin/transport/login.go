package transport

import (
	"N1ktarchik/Bar-Manager/internal/core/transport/request"
	"N1ktarchik/Bar-Manager/internal/core/transport/response"
	"log/slog"
	"net/http"
	"time"
)

func (h *BarAdminHandlerHTTP) LoginAdmin(w http.ResponseWriter, r *http.Request) {
	h.log.Debug("new request POST /api/auth/login",
		slog.Any("time:", time.Now().UTC().Format("2006-01-02 15:04:05")))

	userData := &authDTO{}

	if err := request.DecodeAndValidate(r, userData); err != nil {
		h.log.Debug("parse request error", slog.Any("err", err), slog.Any("Body:", r.Body))

		response.RespondWithError(w, err)
		return
	}

	jwt, err := h.authService.CreateJWT(userData.Password)
	if err != nil {
		response.RespondWithError(w, err)
		return
	}

	h.log.Info("The new admin has been authorized",
		slog.Any("time:", time.Now().UTC().Format("2006-01-02 15:04:05")))

	response.RespondWithJWT(w, http.StatusOK, jwt)
}
