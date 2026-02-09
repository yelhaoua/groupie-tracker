package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetJson(URL string, data any) error {
	// fetching the data from the url
	res, err := http.Get(URL)
	fmt.Println(res , err)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	return json.NewDecoder(res.Body).Decode(data)
}
