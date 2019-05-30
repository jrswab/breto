package wttr

import (
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"strings"
	"time"
)

func Local(cWttr chan string) {
	var passed, hour float64
	start := time.Now() // to determine seconds passed
	ticker := time.NewTicker(time.Second)

	for range ticker.C {
		passed = time.Since(start).Seconds() // total seconds passed
		hour = math.Floor(math.Remainder(passed, 3600))
		if passed < 10 || hour == 0 {
			// get temp(%t) and wind direction/speed (%w)
			// for exact location add postal code - wttr.in/~15222?format...
			// for more wttr options see https://wttr.in/:help
			resp, err := http.Get("https://wttr.in/?format=%t+%w")
			if err != nil {
				errMessage := "wttr connection issue"
				cWttr <- errMessage
			}

			bodyData, _ := ioutil.ReadAll(resp.Body)
			weather := fmt.Sprintf("%s | ",
				strings.TrimSpace(string(bodyData))) // convert responce to string for return
			cWttr <- weather
		}
	}
}
