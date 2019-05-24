package wttr

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func Local() string {
	resp, err := http.Get("https://wttr.in/?format=%t+%w")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	bodyData, _ := ioutil.ReadAll(resp.Body)
	return strings.TrimSpace(string(bodyData))
}
