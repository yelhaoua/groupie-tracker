package handlers

import (
	"net/http"
	"os"
)

/*
	This section contains:
	1. Logic for handling forbidden access
	2. Serving static files
	
*/


func HandleForbiden(w http.ResponseWriter, r *http.Request) {
	// get the pphat of the staticse files
	info, err := os.Stat(r.URL.Path[1:])
	// check if the file not existe
	if err != nil {
		HandlerErr(w, "Page Not Found", http.StatusNotFound)
		return
	}
	// cheak if is directory
	if info.IsDir() {
		HandlerErr(w, "Forbidden", http.StatusForbidden)
		return
	}
	// serving all statics files
	http.ServeFile(w, r, r.URL.Path[1:])
}
