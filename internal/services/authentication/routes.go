package authentication

import (
	"github.com/ocrosby/auth-nexus/internal/services/authentication/handlers"
	"net/http"
)

func SetupRoutes() *http.ServeMux {
	router := http.NewServeMux()

	// Authenticate a user and issue tokens
	router.HandleFunc("POST /login", handlers.LoginHandler)

	// Invalidate tokens.
	router.HandleFunc("POST /logout", handlers.LogoutHandler)

	// Refresh access tokens using refresh tokens
	router.HandleFunc("POST /token/refresh", handlers.TokenRefreshHandler)

	return router
}
