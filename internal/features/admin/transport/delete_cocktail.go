package transport

import (
	"N1ktarchik/Bar-Manager/internal/core/transport/response"
	"log/slog"
	"net/http"
	"time"
)

func (h *BarAdminHandlerHTTP) DeleteCocktail(w http.ResponseWriter, r *http.Request) {
	h.log.Debug("new request DELETE /api/cocktails/{id}",
		slog.Any("time:", time.Now().UTC().Format("2006-01-02 15:04:05")))

	id := r.PathValue("id")

	if err := h.barService.DeleteCocktail(r.Context(), id); err != nil {
		response.RespondWithError(w, err)
		return
	}

	h.log.Info("The cocktail has been deleted",
		slog.Any("ID", id),
		slog.Any("time:", time.Now().UTC().Format("2006-01-02 15:04:05")))

	response.RespondWithJSON(w, http.StatusNoContent, "")
}
