package blocks

import (
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"strings"
	"time"
)

func Wttr(cWttr chan string, eWttr chan error) {
	var passed, hour float64
	start := time.Now() // to determine seconds passed
	ticker := time.NewTicker(time.Second)

	for range ticker.C {
		passed = time.Since(start).Seconds() // total seconds passed
		hour = math.Floor(math.Remainder(passed, 3600))

		if passed < 10 || hour == 0 {
			// for options see https://wttr.in/:help
			resp, err := http.Get("https://wttr.in/?format=%t+%w")
			if err != nil {
				eWttr <- err
			}

			bodyData, _ := ioutil.ReadAll(resp.Body)
			// convert responce to string for go channel
			weather := fmt.Sprintf("%s | ",
				strings.TrimSpace(string(bodyData)))
			resp.Body.Close()
			cWttr <- weather
		}
	}
}
