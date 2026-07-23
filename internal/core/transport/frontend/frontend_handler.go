package frontend

import (
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
)

type FrontendHandler struct {
	staticDir string
}

func NewFrontendHandler(staticDir string) *FrontendHandler {
	return &FrontendHandler{
		staticDir: staticDir,
	}
}

func (h *FrontendHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/", h.serveIndex).Methods(http.MethodGet)
	router.HandleFunc("/admin", h.serveAdmin).Methods(http.MethodGet)
}

func (h *FrontendHandler) serveIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, filepath.Join(h.staticDir, "index.html"))
}

func (h *FrontendHandler) serveAdmin(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, filepath.Join(h.staticDir, "admin.html"))
}
