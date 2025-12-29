package main

import (
	"net/http"

	"groupie/handlers"
)

func main() {
	http.HandleFunc("/statics/", handlers.HandleForbiden)
	http.HandleFunc("/", handlers.HandleHome)
	http.HandleFunc("/artist/{id}", handlers.HandleInfo)
	http.HandleFunc("/artist/", handlers.HandleArtist)

	http.ListenAndServe(":8080", nil)
}
