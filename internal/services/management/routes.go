package management

import "net/http"

func SetupRoutes() *http.ServeMux {
	router := http.NewServeMux()

	// Example handler
	router.HandleFunc("GET /whatever", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Example handler"))
	})

	return router
}
