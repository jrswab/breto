package wttr

import (
	"io/ioutil"
	"net/http"
	"strings"
)

func Local() *string {
	// get temp(%t) and wind direction/speed (%w)
	// for exact location add postal code - wttr.in/~15222?format...
	// for more wttr options see https://wttr.in/:help
	resp, err := http.Get("https://wttr.in/?format=%t+%w")
	if err != nil {
		errMessage := "wttr connection issue"
		return &errMessage
	}
	defer resp.Body.Close() // close http request

	// convert responce to string for return
	bodyData, _ := ioutil.ReadAll(resp.Body)
	weather := strings.TrimSpace(string(bodyData))
	return &weather
}
