package user

import "net/http"

func SetupRoutes() *http.ServeMux {
	router := http.NewServeMux()

	// Register a new user.
	router.HandleFunc("POST /register", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Register a new user"))
	})

	// Get user details.
	router.HandleFunc("GET /users/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Get user details"))
	})

	// Update user details.
	router.HandleFunc("PUT /users/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Update user details"))
	})

	// Delete a user.
	router.HandleFunc("DELETE /users/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Delete a user"))
	})

	return router
}
