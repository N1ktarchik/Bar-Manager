package transport

import (
	"N1ktarchik/Bar-Manager/internal/core/transport/request"
	"N1ktarchik/Bar-Manager/internal/core/transport/response"
	"log/slog"
	"net/http"
	"time"
)

func (h *BarAdminHandlerHTTP) AddCocktailHandler(w http.ResponseWriter, r *http.Request) {
	h.log.Debug("new request POST /api/cocktails",
		slog.Any("time:", time.Now().UTC().Format("2006-01-02 15:04:05")))

	userData := &cocktailDTO{}

	if err := request.DecodeAndValidate(r, userData); err != nil {
		h.log.Debug("parse request error", slog.Any("err", err), slog.Any("Body:", r.Body))

		response.RespondWithError(w, err)
		return
	}

	cocktail, err := h.service.AddCocktail(r.Context(), userData.ToDomain())
	if err != nil {
		response.RespondWithError(w, err)
		return
	}

	h.log.Info("A new cocktail has been added", slog.Int("ID", cocktail.Id),
		slog.Any("time:", time.Now().UTC().Format("2006-01-02 15:04:05")))

	response.RespondWithJSON(w, http.StatusCreated, cocktail)

}
