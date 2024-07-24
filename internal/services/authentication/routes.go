package router

import "net/http"

func SetupRoutes() *http.ServeMux {
	router := http.NewServeMux()

	// Authenticate a user and issue tokens
	router.HandleFunc("POST /login", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Login service"))
	})

	// Invalidate tokens.
	router.HandleFunc("POST /logout", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Logout service"))
	})

	// Refresh access tokens using refresh tokens
	router.HandleFunc("POST /token/refresh", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Token refresh service"))
	})

	return router
}
