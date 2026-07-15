package transport

import (
	"N1ktarchik/Bar-Manager/internal/core/transport/request"
	"N1ktarchik/Bar-Manager/internal/core/transport/response"
	"log/slog"
	"net/http"
)

func (h *BarAdminHandlerHTTP) AddCocktailHandler(w http.ResponseWriter, r *http.Request) {
	h.log.Debug("new request POST /api/cocktails")
	userData := &cocktailDTO{}

	if err := request.DecodeAndValidate(r, userData); err != nil {
		h.log.Debug("parse request error", slog.Any("err", err), slog.Any("Body:", r.Body))

		response.RespondWithError(w, err)
		return
	}

	cocktail, err := h.barService.AddCocktail(r.Context(), userData.ToDomain())
	if err != nil {
		response.RespondWithError(w, err)
		return
	}

	h.log.Info("A new cocktail has been added", slog.Int("ID", cocktail.Id))

	response.RespondWithJSON(w, http.StatusCreated, cocktail)

}
