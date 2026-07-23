package transport

import (
	"N1ktarchik/Bar-Manager/internal/core/transport/request"
	"N1ktarchik/Bar-Manager/internal/core/transport/response"
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
)

func (h *BarAdminHandlerHTTP) UpdatePriceHandler(w http.ResponseWriter, r *http.Request) {
	h.log.Debug("new request PATCH /api/cocktails/{id}/price")

	userData := &cocktailDTO{}

	if err := request.DecodeAndValidate(r, userData); err != nil {
		h.log.Debug("parse request error", slog.Any("err", err), slog.Any("Body:", r.Body))

		response.RespondWithError(w, err)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]

	updatedCocktail, err := h.barService.UpdatePrice(r.Context(), id, userData.Price)
	if err != nil {
		response.RespondWithError(w, err)
		return
	}

	h.log.Info("The price of a cocktail has been updated",
		slog.Int("ID", updatedCocktail.Id),
		slog.Int("new price", updatedCocktail.Price))

	response.RespondWithJSON(w, http.StatusOK, updatedCocktail)

}
