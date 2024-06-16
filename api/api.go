package main

import (
	"log"
	"net/http"
)

// Let me create an api from scratch to explore Go http package
// GO Version: 1.22.3

// APIServer is a struct for every http server I will create
type APIServer struct {
	addr string
}

// NewAPIServer is a factory for my Server.
func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

// Run method starts the server
func (s *APIServer) Run() error {
	router := http.NewServeMux()

	// We return a path value in the request url in this endpoint
	// @GET /users/{userID}
	router.HandleFunc("GET /users/{UserID}", func(w http.ResponseWriter, r *http.Request) {

		// Get PathValue
		userID := r.PathValue("UserID")

		// Response Status: 200(OK)
		w.WriteHeader(http.StatusOK)

		// Return UserID
		w.Write([]byte(userID))

	})

	// Implementing sub-routing
	v1 := http.NewServeMux()
	v1.Handle("/api/v1/", http.StripPrefix("/api/v1/", router))

	middlewareChain := MiddleWareChain(
		RequestLoggerMiddleware,
		RequireAuthMiddleWare,
	)

	// Set server parameters
	server := http.Server{
		Addr: s.addr,
		// Chaining middleware
		Handler: middlewareChain(router),
	}

	log.Printf("Server is listening at %s", s.addr)

	return server.ListenAndServe()
}

// RequestLoggerMiddleware logs a request to the console.
// It returns the request Method and the URL path.
// Example:2024/06/16 16:07:54 METHOD: @GET
// Path: /users/01
func RequestLoggerMiddleware(next http.Handler) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("METHOD: @%s\nPath: %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	}
}

// RequireAuthMiddleWare protects routes
func RequireAuthMiddleWare(next http.Handler) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		// check if user is authenticated
		token := r.Header.Get("Authorization")

		// Mock a request token
		// Request should look like:
		// $ curl -H "Authorization: Bearer token" http://localhost:4032/users/01
		if token != "Bearer token" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
}

// MiddleWare type
type MiddleWare func(handler http.Handler) http.HandlerFunc

// MiddleWareChain chains middlewares and wraps our routes
// It takes args of middlewares and return a handlerFunc
func MiddleWareChain(middlewares ...MiddleWare) MiddleWare {
	return func(next http.Handler) http.HandlerFunc {

		for i := len(middlewares) - 1; i >= 0; i-- {
			next = middlewares[i](next)
		}
		return next.ServeHTTP
	}
}
