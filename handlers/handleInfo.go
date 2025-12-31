package handlers

import (
	"net/http"
	"strconv"
)

/*
	This section contains:
	1. Get the artits id from url
	2. fetch all artist data name, loctions, dates ...
	3. Send the final data to page of artist info

*/

func HandleInfo(w http.ResponseWriter, r *http.Request) {
	var all_info []Artist
	var LocatonsStruct LocationStruct
	var DatesStruct DatesStruct
	var RelationsStruct RelationStruct
	var ArtistData Artist

	// check which methode usde
	if r.Method != http.MethodGet {
		HandlerErr(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	// converting the ID
	artist_id, num_error := strconv.Atoi(r.PathValue("id"))
	if num_error != nil {
		HandlerErr(w, "Bad Request", http.StatusBadRequest)
		return
	}
	// fetch all url and check the msg errore
	artistError := GetJson(UrlArtist, &all_info)

	// find the artist with same id
	for _, val := range all_info {
		if val.Id == artist_id {
			ArtistData = val
			break
		}
	}

	locationError := GetJson(ArtistData.Locations, &LocatonsStruct)
	datesError := GetJson(ArtistData.ConcertDates, &DatesStruct)
	relationrror := GetJson(ArtistData.Relations, &RelationsStruct)
	if artistError != nil || locationError != nil || datesError != nil || relationrror != nil {
		HandlerErr(w, "Not Found", http.StatusNotFound)
		return
	}

	// save the all data in one page to serve it
	data := AllData{
		Artist:   ArtistData,
		Location: LocatonsStruct,
		Dates:    DatesStruct,
		Relation: RelationsStruct,
	}

	// execute the artistInfo template
	error_page := Temp.ExecuteTemplate(w, "artistInfo.html", data)
	if error_page != nil {
		HandlerErr(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
