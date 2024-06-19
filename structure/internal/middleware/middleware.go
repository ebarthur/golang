package middleware

import (
	"fmt"
	"net/http"
)

// RequestLoggerMiddleware manually logs every request in the terminal
func RequestLoggerMiddleware(next http.Handler) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		fmt.Printf("METHOD: %s\nURL PATH: %s\n", r.Method, r.URL.Path)

		next.ServeHTTP(w, r)
	}
}

// RequestAuthMiddleware mocks authentication and may be used to protect routes
func RequestAuthMiddleware(next http.Handler) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		// TODO: implement go-jwt
		token := r.Header.Get("Authorization")

		if pass := "Bearer token"; token != pass {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
}

// Middleware is a function type that takes a handler and returns a handler function.
// In Go, functions are values*
type Middleware func(http.Handler) http.HandlerFunc

// Middlewares is a function that chains middleware functions
//
// So basically, we pass middleware functions that return http.Handler and finally
// return a handler that returns a handler function.
// It's not that hard. Wrap your head around it.
func Middlewares(middleware ...Middleware) Middleware {

	return func(next http.Handler) http.HandlerFunc {

		for i := len(middleware) - 1; i >= 0; i-- {
			next = middleware[i](next)
		}

		return next.ServeHTTP
	}

}
