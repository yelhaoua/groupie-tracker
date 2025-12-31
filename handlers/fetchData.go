package handlers

import (
	"encoding/json"
	"net/http"
)
/*
	This section contains:
	1. Logic for fetching the url data and save it

*/


func GetJson(URL string, data any) error {
	// fetching the data from the url
	res, err := http.Get(URL)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	return json.NewDecoder(res.Body).Decode(data)
}
