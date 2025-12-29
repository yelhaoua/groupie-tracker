package handlers

import (
	"net/http"
	"text/template"
)

var Temp, ErrParse = template.ParseGlob("./templates/*.html")

func HandleHome(w http.ResponseWriter, r *http.Request) {
	var data []Artist
	if r.URL.Path == "/" {
		if r.Method != http.MethodGet {
			HandlerErr(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		url := "https://groupietrackers.herokuapp.com/api/artists"
		err := GetJson(url, &data)
		if err != nil {
			HandlerErr(w, "Not Found", http.StatusNotFound)
			return
		}
		Temp.ExecuteTemplate(w, "index.html", data)
	} else {
		HandlerErr(w, "Page Not Found", http.StatusNotFound)
		return
	}
}
