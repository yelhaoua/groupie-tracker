package handlers

import (
	"net/http"
	"html/template"
)

/*
	This section contains:
	1. Check if valid apath of home and method of request
	2. Fetch Artists data to show on home page
	3. Execute Index with Artitst data

*/

// global declarations
var Temp, ErrParse = template.ParseGlob("./templates/*.html")
var UrlArtist = "https://groupietrackers.herokuapp.com/api/artists"

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
		
		err := GetJson(UrlArtist, &data)
		if err != nil {
			HandlerErr(w, "Not Found", http.StatusNotFound)
			return
		}
		error_page := Temp.ExecuteTemplate(w, "index.html", data)
		if error_page != nil {
			HandlerErr(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	} else {
		HandlerErr(w, "Page Not Found", http.StatusNotFound)
		return
	}
}
