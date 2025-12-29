package handlers

import (
	"net/http"
	"os"
)

func HandleForbiden(w http.ResponseWriter, r *http.Request) {
	info, err := os.Stat(r.URL.Path[1:])
	if err != nil {
		HandlerErr(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if info.IsDir() {
		HandlerErr(w, "Forbidden", http.StatusForbidden)
		return
	}

	http.ServeFile(w, r, r.URL.Path[1:])
}
