package transport

import (
	"N1ktarchik/Bar-Manager/internal/core/transport/request"
	"N1ktarchik/Bar-Manager/internal/core/transport/response"
	"log/slog"
	"net/http"
	"time"
)

func (h *BarAdminHandlerHTTP) UpdatePriceHandler(w http.ResponseWriter, r *http.Request) {
	h.log.Debug("new request PATCH /api/cocktails/{id}/price",
		slog.Any("time:", time.Now().UTC().Format("2006-01-02 15:04:05")))

	userData := &cocktailDTO{}

	if err := request.DecodeAndValidate(r, userData); err != nil {
		h.log.Debug("parse request error", slog.Any("err", err), slog.Any("Body:", r.Body))

		response.RespondWithError(w, err)
		return
	}

	id := r.PathValue("id")

	updatedCocktail, err := h.service.UpdatePrice(r.Context(), id, userData.Price)
	if err != nil {
		response.RespondWithError(w, err)
		return
	}

	h.log.Info("The price of a cocktail has been updated", slog.Int("ID", updatedCocktail.Id),
		slog.Int("new price", updatedCocktail.Price),
		slog.Any("time:", time.Now().UTC().Format("2006-01-02 15:04:05")))

	response.RespondWithJSON(w, http.StatusOK, updatedCocktail)

}
