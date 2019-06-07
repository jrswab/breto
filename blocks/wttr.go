package blocks

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"strings"
	"time"
)

// Wttr gets the weather of the computer's general location.
// Specify city or area code: "wttr.in/~15222" or "wttr.in/~Pittsuburgh"
func Wttr(cWttr chan string, eWttr chan error) {
	var passed, hour float64
	var data string

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
			if strings.Contains(string(bodyData), "error") { // wttr.in displays a webpage on server error
				eWttr <- errors.New("wttr.in overloaded") // display this on wttr.in server error
			}
			// convert responce to string for go channel
			data = string(bodyData)
			weather := fmt.Sprintf("%s |", strings.TrimSpace(data))
			resp.Body.Close()
			cWttr <- weather
		}
	}
}
