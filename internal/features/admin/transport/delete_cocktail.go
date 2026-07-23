package transport

import (
	"N1ktarchik/Bar-Manager/internal/core/transport/response"
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
)

func (h *BarAdminHandlerHTTP) DeleteCocktail(w http.ResponseWriter, r *http.Request) {
	h.log.Debug("new request DELETE /api/cocktails/{id}")
	vars := mux.Vars(r)
	id := vars["id"]

	if err := h.barService.DeleteCocktail(r.Context(), id); err != nil {
		response.RespondWithError(w, err)
		return
	}

	h.log.Info("The cocktail has been deleted",
		slog.Any("ID", id))

	response.RespondWithJSON(w, http.StatusNoContent, "")
}
