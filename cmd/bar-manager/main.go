package main

import (
	"N1ktarchik/Bar-Manager/internal/core/logger"
	postgres "N1ktarchik/Bar-Manager/internal/core/repository"
	"N1ktarchik/Bar-Manager/internal/core/transport/frontend"
	"N1ktarchik/Bar-Manager/internal/core/transport/middleware"
	"N1ktarchik/Bar-Manager/internal/core/transport/server"
	admin_repo "N1ktarchik/Bar-Manager/internal/features/admin/repository"
	admin_serv "N1ktarchik/Bar-Manager/internal/features/admin/service"
	admin_http "N1ktarchik/Bar-Manager/internal/features/admin/transport"
	"N1ktarchik/Bar-Manager/internal/features/auth"
	client_repo "N1ktarchik/Bar-Manager/internal/features/client/repository"
	client_serv "N1ktarchik/Bar-Manager/internal/features/client/service"
	client_http "N1ktarchik/Bar-Manager/internal/features/client/transport"
	"context"
	"log/slog"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	log := logger.Setup()
	slog.SetDefault(log)

	log.Info("starting bar-manager")

	if err := godotenv.Load(); err != nil {
		log.Error(".env file not found")
		return
	}

	connStr := postgres.GetPostgresValues()
	config := postgres.NewPostgresConfig(connStr, 25, 5, 30*time.Minute, 5*time.Minute, 1*time.Minute)

	pool, err := postgres.CreatePool(context.Background(), config, log)
	if err != nil {
		log.Error("failed to initialize database pool", slog.Any("err", err))
		return
	}
	defer pool.Close()
	log.Info("postgres connection pool established")

	ADMIN_PASS := os.Getenv("ADMIN_PASSWORD")
	if ADMIN_PASS == "" {
		log.Error("admin password is not set in .env")
		return
	}

	SECRET_KEY := os.Getenv("SECRET_KEY")
	if SECRET_KEY == "" {
		log.Error("secret key is not set in .env")
		return
	}

	authService := auth.NewJWTService([]byte(SECRET_KEY), ADMIN_PASS, log)

	clientRepository := client_repo.NewBarClientRepository(pool, log)
	clientService := client_serv.NewBarClientService(clientRepository)
	clientTransport := client_http.NewBarClientTransportHTTP(clientService, log)

	adminRepository := admin_repo.NewBarAdminRepository(pool, log)
	adminService := admin_serv.NewBarAdminService(adminRepository, log)
	adminTransport := admin_http.NewBarAdminTransportHTTP(adminService, authService, log)

	mw := middleware.NewMiddleware(authService, log)
	srv := server.NewServer(log)
	r := srv.Router

	frontendHandler := frontend.NewFrontendHandler("./web")
	frontendHandler.RegisterRoutes(r)

	api := r.PathPrefix("/api").Subrouter()

	api.HandleFunc("/cocktails", clientTransport.GetCocktailsHandler).Methods("GET")
	api.HandleFunc("/auth/login", adminTransport.LoginAdmin).Methods("POST")

	secureAPI := api.PathPrefix("").Subrouter()
	secureAPI.Use(mw.AuthMiddleware)

	secureAPI.HandleFunc("/cocktails", adminTransport.AddCocktailHandler).Methods("POST")
	secureAPI.HandleFunc("/cocktails/{id}/price", adminTransport.UpdatePriceHandler).Methods("PATCH")
	secureAPI.HandleFunc("/cocktails/{id}", adminTransport.DeleteCocktail).Methods("DELETE")

	log.Info("all services initialized, transport starting", slog.String("addr", ":8080"))
	if err := srv.Run(":8080"); err != nil {
		log.Error("server crashed", slog.Any("err", err))
	}
}
