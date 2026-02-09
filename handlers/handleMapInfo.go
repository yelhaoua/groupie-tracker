package handlers

import (
	"net/http"
	"strings"
)

func HandleMap(w http.ResponseWriter, r *http.Request) {
	parames := r.PathValue("name")
	parames = strings.ReplaceAll(parames, "_", "+")
	parames = strings.ReplaceAll(parames, "-", ",+")
	Temp.ExecuteTemplate(w, "map.html", parames)
}
