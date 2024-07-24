package router

import "net/http"

func SetupRoutes() *http.ServeMux {
	router := http.NewServeMux()

	// Issue tokens
	router.HandleFunc("POST /token", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Issue tokens"))
	})

	// Validate Tokens
	router.HandleFunc("POST /token/validate", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Validate tokens"))
	})

	// Revoke Tokens
	router.HandleFunc("POST /token/revoke", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Revoke tokens"))
	})

	return router
}
