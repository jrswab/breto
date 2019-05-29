package wttr

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func Local(c chan string) {
	weather := ""
	for i := 0; i < 3636; {
		// get temp(%t) and wind direction/speed (%w)
		// for exact location add postal code - wttr.in/~15222?format...
		// for more wttr options see https://wttr.in/:help
		if i == 0 || i == 3600 {
			resp, err := http.Get("https://wttr.in/?format=%t+%w")
			if err != nil {
				errMessage := "wttr connection issue"
				c <- errMessage
			}
			defer resp.Body.Close() // close http request

			// convert responce to string for return
			bodyData, _ := ioutil.ReadAll(resp.Body)
			weather = fmt.Sprintf("%s | ", strings.TrimSpace(string(bodyData)))
		}
		c <- weather
		if i == 3600 {
			i = 0
		} else {
			i++
		}
		time.Sleep(1 * time.Second)
	}
}
