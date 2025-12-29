package main

import (
	"fmt"
	"net/http"

	"groupie/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/statics/", handlers.HandleForbiden)
	mux.HandleFunc("/", handlers.HandleHome)
	mux.HandleFunc("/artist/{id}", handlers.HandleInfo)
	mux.HandleFunc("/artistes", handlers.HandleArtist)
	fmt.Printf("Server listning on : http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}
