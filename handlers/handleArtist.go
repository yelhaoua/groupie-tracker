package handlers

import (
	"net/http"
)

type Artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

var ArrArtist []Artist

func HandleArtist(w http.ResponseWriter, r *http.Request) {
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
	url := "https://groupietrackers.herokuapp.com/api/artists"
	// fetching the data from url
	err := GetJson(url, &ArrArtist)
	// chek if the fetch return errore
	if err != nil {
		HandlerErr(w, "fetching Errore", http.StatusNotFound)
		return
	}
	// execute the template 
	Temp.ExecuteTemplate(w, "allartist.html", ArrArtist)
}
