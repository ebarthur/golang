package routes

import (
	"github.com/ebarthur/go-app-structure/internal/handlers"
	mid "github.com/ebarthur/go-app-structure/internal/middleware"
	"net/http"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	// Main handler for /api/v1/ routes
	v1 := http.NewServeMux()

	// Handler for protected routes
	v1Auth := http.NewServeMux()

	v1.Handle("/ping", handlers.HandlePing())
	v1Auth.Handle("/fact", handlers.HandleCatFact())
	v1Auth.Handle("/data", handlers.HandleData())

	// Middleware wrap for routes
	UnprotectedRoutes := mid.Middlewares(mid.RequestLoggerMiddleware)
	ProtectedRoutes := mid.Middlewares(mid.RequestLoggerMiddleware, mid.RequestAuthMiddleware)

	// Use StripPrefix to strip /api/v1 and forward to v1 handler
	mux.Handle("/api/", http.StripPrefix("/api", UnprotectedRoutes(v1)))

	// Use StripPrefix to strip /api/v1 and forward to v1_auth handler
	// This handler protects routes that require authentication
	mux.Handle("/api/v1/", http.StripPrefix("/api/v1", ProtectedRoutes(v1Auth)))

	return mux
}
