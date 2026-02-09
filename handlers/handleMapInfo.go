package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
)

func HandleMap(w http.ResponseWriter, r *http.Request) {
	var Maps GeoResponse

	name := r.PathValue("name")
	name = strings.ReplaceAll(name, "_", "+")
	name = strings.ReplaceAll(name, "-", "+")

	url := "https://address-from-to-latitude-longitude.p.rapidapi.com/geolocationapi?address=" + name

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-key", "d55037e0b9msh675e1a09cd24677p171435jsne9bf014676c9")
	req.Header.Add("x-rapidapi-host", "address-from-to-latitude-longitude.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)
	err := json.NewDecoder(res.Body).Decode(&Maps)
	if err != nil {
		HandlerErr(w, "Map  Not found", 403)
		return
	}
	defer res.Body.Close()

	Temp.ExecuteTemplate(w, "map.html", Maps.Results[0])
}
