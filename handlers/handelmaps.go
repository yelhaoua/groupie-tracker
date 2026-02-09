package handlers

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type Geo struct {
	Features []GeoFeature `json:"features"`
}

type GeoFeature struct {
	Geometry GeoGeometry `json:"geometry"`
}

type GeoGeometry struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

func CleanQuery(str string) string {
	final := ""
	for _, v := range str {
		if v == '-' {
			final += ", "
		} else if v == '_' {
			final += " "

		} else {

			final += string(v)
		}
	}
	return final
}

func GetGeo(w http.ResponseWriter, r *http.Request) {
	var geo Geo

	if r.Method != http.MethodGet {
		HandlerErr(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	query := CleanQuery(r.URL.Query().Get("q"))
	if query == "" {
		query = "Los Angeles, usa"
	}

	apiURL := "https://api.mapbox.com/search/geocode/v6/forward?q=" + url.QueryEscape(query) + "&access_token=pk.eyJ1IjoiaGFpdGJlbmFsIiwiYSI6ImNtbGZldXluMjAxNnMzbHM4cGhuY2M3dzgifQ.8TamijQFnTPdolvoT55Ttg"
	err := GetJson(apiURL, &geo)
	if err != nil {
		HandlerErr(w, "fetching Errore", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(geo.Features); err != nil {
		HandlerErr(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
