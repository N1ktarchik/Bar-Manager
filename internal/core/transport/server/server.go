package server

import (
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	Router *mux.Router
	log    *slog.Logger
}

func NewServer(log *slog.Logger) *Server {
	return &Server{
		Router: mux.NewRouter(),
		log:    log,
	}
}

func (s *Server) Run(addr string) error {
	s.log.Info("starting server", slog.String("addr", addr))

	err := http.ListenAndServe(addr, s.Router)
	if err != nil {
		s.log.Error("server failed to start or crashed",
			slog.String("addr", addr),
			slog.Any("err", err),
		)
		return err
	}

	return nil
}
