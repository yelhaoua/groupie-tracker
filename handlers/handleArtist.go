package handlers

import (
	"net/http"
)

/*
	This section contains:
	1. Logic for handling Artist
*/

func HandleArtist(w http.ResponseWriter, r *http.Request) {
	var ArrArtist []Artist

	// check if parsing globale return errore
	if ErrParse != nil {
		HandlerErr(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// check which methode used
	if r.Method != http.MethodGet {
		HandlerErr(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// fetching the data from url
	err := GetJson(UrlArtist, &ArrArtist)
	// chek if the fetch return errore
	if err != nil {
		HandlerErr(w, "fetching Errore", http.StatusNotFound)
		return
	}
	// execute the template
	error_page := Temp.ExecuteTemplate(w, "allartist.html", ArrArtist)
	if error_page != nil {
		HandlerErr(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
