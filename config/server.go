package config

import (
	"log"
	"net/http"

	"github.com/ChileKasoka/mis/cmd/api"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

type Server struct {
	router *chi.Mux
}

func MiddlewareLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func NewServer(
	appointmentHandler *api.AppointmentHandler,
	userHandler *api.UserHandler) *Server {
	// Initialize a new Chi router
	router := chi.NewRouter()

	// Configure CORS options
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// Create a sub-router for versioned routes
	v1Router := chi.NewRouter()
	v1Router.Use(MiddlewareLogger)

	// Register appointment routes
	appointmentHandler.RegisterRoutes(v1Router)
	userHandler.RegisterRoutes(v1Router)

	// Mount the v1 API to the main router
	router.Mount("/v1", v1Router)

	return &Server{
		router: router,
	}
}

func (server *Server) Start(address string) error {
	srv := &http.Server{
		Addr:    address,
		Handler: server.router,
	}

	log.Printf("Starting server on %s", address)
	if err := srv.ListenAndServe(); err != nil {
		log.Printf("Server error: %v", err)
		return err
	}
	return nil
}
