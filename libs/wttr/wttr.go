package wttr

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func Local(cWttr chan string) {
	// get temp(%t) and wind direction/speed (%w)
	// for exact location add postal code - wttr.in/~15222?format...
	// for more wttr options see https://wttr.in/:help
	resp, err := http.Get("https://wttr.in/?format=%t+%w")
	if err != nil {
		errMessage := "wttr connection issue"
		cWttr <- errMessage
	}

	// convert responce to string for return
	bodyData, _ := ioutil.ReadAll(resp.Body)
	weather := fmt.Sprintf("%s | ", strings.TrimSpace(string(bodyData)))
	cWttr <- weather
	resp.Body.Close() // close http request
}
