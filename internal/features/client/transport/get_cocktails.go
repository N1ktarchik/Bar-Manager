package transport

import (
	"N1ktarchik/Bar-Manager/internal/core/transport/response"
	"log/slog"
	"net/http"
)

func (h *BarClientHandlerHTTP) GetCocktailsHandler(w http.ResponseWriter, r *http.Request) {
	h.log.Debug("new request GET /api/cocktails")

	cocktails, err := h.service.GetCocktails(r.Context())

	if err != nil {
		h.log.Error("error to get cocktails", slog.Any("err", err))
		response.RespondWithError(w, err)
		return
	}

	h.log.Debug("the cocktails were received successfully")
	response.RespondWithArray(w, http.StatusOK, "", cocktails)
}
