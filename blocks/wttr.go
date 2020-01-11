package blocks

import (
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"strings"
	"time"
)

// Wttr gets the weather of the computer's general location.
// Specify city or area code: "wttr.in/~15222" or "wttr.in/~Pittsburgh"
func Wttr(cWttr chan string, eWttr chan error) {
	var passed, hour float64
	var data, weather string

	start := time.Now() // to determine seconds passed
	ticker := time.NewTicker(time.Second)

	for range ticker.C {
		passed = time.Since(start).Seconds() // total seconds passed
		hour = math.Floor(math.Remainder(passed, 3600))

		if passed < 10 || hour == 0 {
			// for options see https://wttr.in/:help
			resp, err := http.Get("https://wttr.in/?format=%t")
			if err != nil {
				eWttr <- err
			}

			bodyData, _ := ioutil.ReadAll(resp.Body)
			resp.Body.Close()
			data = string(bodyData)

			// wttr.in sends an html string during the server error.
			// Several attempts to catch the exact error proved difficult
			// so now we catch the correct outupt instead.
			if strings.Contains(data, "+") || strings.Contains(data, "-") {
				weather = fmt.Sprintf("%s", strings.TrimSpace(data))
				cWttr <- weather
			} else {
				eWttr <- fmt.Errorf("Expected temp, got: %s", data)
			}
		}
	}
}
