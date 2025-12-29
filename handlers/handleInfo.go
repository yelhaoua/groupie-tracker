package handlers

import (
	"net/http"
	"strconv"
	"strings"
)

type LocationStruct struct {
	Id       int      `json:"id"`
	Location []string `json:"locations"`
}
type DatesStruct struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}

type RelationStruct struct {
	Id             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type AllData struct {
	Artist   Artist
	Location LocationStruct
	Dates    DatesStruct
	Relation RelationStruct
}

func HandleInfo(w http.ResponseWriter, r *http.Request) {
	var allinfo []Artist
	var Locatonsstruct LocationStruct
	var Datesstruct DatesStruct
	var Relationstruct RelationStruct
	// check which methode usde
	if r.Method != http.MethodGet {
		HandlerErr(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	// converting the ID
	num, nErr := strconv.Atoi(strings.TrimSuffix(r.URL.Path[len("/artist/"):], "/"))
	id := strings.TrimSuffix(r.URL.Path[len("/artist/"):], "/")
	if nErr != nil {
		HandlerErr(w, "Page Not Found", http.StatusNotFound)
		return
	}
	// fetch all url and check the msg errore
	Aerr := GetJson("https://groupietrackers.herokuapp.com/api/artists", &allinfo)
	if r.URL.Path != "/artist/"+id {
		HandlerErr(w, "Page Not Found", http.StatusNotFound)
		return
	}
	locations := "https://groupietrackers.herokuapp.com/api/locations/" + id
	Lerr := GetJson(locations, &Locatonsstruct)
	dates := "https://groupietrackers.herokuapp.com/api/dates/" + id
	Derr := GetJson(dates, &Datesstruct)
	relation := "https://groupietrackers.herokuapp.com/api/relation/" + id
	Rerr := GetJson(relation, &Relationstruct)
	if Aerr != nil || Lerr != nil || Derr != nil || Rerr != nil {
		HandlerErr(w, "Not Found", http.StatusNotFound)
		return
	}
	var ArtistData Artist

	// finde the asrtiste based on the ID
	for _, val := range allinfo {
		if val.Id == num {
			ArtistData = val
			break
		}
	}
	data := AllData{
		Artist:   ArtistData,
		Location: Locatonsstruct,
		Dates:    Datesstruct,
		Relation: Relationstruct,
	}
	// execute the artistInfo template
	Temp.ExecuteTemplate(w, "artistInfo.html", data)
}
