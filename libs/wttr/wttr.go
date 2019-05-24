package wttr

import (
	"io/ioutil"
	"net/http"
	"strings"
)

func Local() string {
	// for more wttr options see https://wttr.in/:help
	resp, err := http.Get("https://wttr.in/?format=%t+%w") // get temp and wind direction/speed
	if err != nil {
		return err.Error()
	}
	defer resp.Body.Close() // close http request

	// convert responce to string for return
	bodyData, _ := ioutil.ReadAll(resp.Body)
	return strings.TrimSpace(string(bodyData))
}
