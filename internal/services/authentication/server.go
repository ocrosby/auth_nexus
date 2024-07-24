package authentication

import (
	"github.com/ocrosby/auth-nexus/internal/authentication/router"
	"net/http"
)

type ServiceInterface interface {
	Run() error
}

type Server struct {
	Router *http.ServeMux
}

func NewServer(
	logger *Logger,
	config *Config,
	commentStore *commentStore,
	anotherStore *anotherStore,
) http.Handler {
	mux := router.SetupRoutes()
	addRoutes(
		mux,
		Logger,
		Config,
		commentStore,
		anotherStore,
	)
	var handler http.Handler = mux
	handler = someMiddleware(handler)
	handler = someMiddleware2(handler)
	handler = someMiddleware3(handler)
	return handler
}

func (s *Server) Run() error {
	return nil
}
