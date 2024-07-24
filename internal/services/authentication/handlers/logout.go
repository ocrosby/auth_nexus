package handlers

import "net/http"

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Logout handler"))
}
