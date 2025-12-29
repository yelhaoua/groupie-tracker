package handlers

import (
	"net/http"
	"strings"
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
	if ErrParse != nil {
		HandlerErr(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	url := "https://groupietrackers.herokuapp.com/api/artists"
	err := GetJson(url, &ArrArtist)
	if err != nil {
		HandlerErr(w, "fetching Errore", http.StatusNotFound)
		return
	}
	if r.URL.Path == "/artist/" {
		Temp.ExecuteTemplate(w, "allartist.html", ArrArtist)
		return
	} else if strings.Contains(r.URL.Path, "/artist/") && r.URL.Path != "/artist/" {
		HandlerErr(w, "Page Not Found", http.StatusNotFound)
		return
	}
	Temp.ExecuteTemplate(w, "index.html", ArrArtist)
}
