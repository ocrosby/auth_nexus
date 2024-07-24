package router

import "net/http"

func SetupRoutes() *http.ServeMux {
	router := http.NewServeMux()

	// Authorize a request based on tokens
	router.HandleFunc("GET /authorize", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Authorize a request based on tokens"))
	})

	// Create a new roles
	router.HandleFunc("POST /roles", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Create new roles"))
	})

	// Create a new permissions
	router.HandleFunc("POST /permissions", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Create new permissions"))
	})

	// Assign roles to users.
	router.HandleFunc("POST /assign-role", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Assign roles to users"))
	})

	// Assign permissions to roles.
	router.HandleFunc("POST /assign-permission", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Assign permissions to roles"))
	})

	return router
}
