package handlers

import (
	"net/http"
	"text/template"
)

var Temp, ErrParse = template.ParseGlob("./templates/*.html")

func HandleHome(w http.ResponseWriter, r *http.Request) {
	var data []Artist
	// cheak if the user in home page 
	if r.URL.Path == "/" {
		// check which method used
		if r.Method != http.MethodGet {
			HandlerErr(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		// fetching the url and check if hase an errore
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
