package handlers

import "net/http"

func TokenRefreshHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Token refresh handler"))
}
