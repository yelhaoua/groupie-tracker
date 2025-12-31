package handlers

import (
	"net/http"
)

/*
	This section contains:
	1. Logic for handling Error code and description
	2. Write the status code and response

*/

func HandlerErr(w http.ResponseWriter, description string, code int) {
	pageErr := info{Code: code, Description: description}
	// write in the header the status code end execute the error page 
	if ErrParse != nil {
		w.WriteHeader(code)
		Temp.ExecuteTemplate(w, "errorePage.html", pageErr)
		return
	}
	w.WriteHeader(code)
	Temp.ExecuteTemplate(w, "errorePage.html", pageErr)
}
