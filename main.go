package main

import (
	"fmt"
	"os/exec"
	"strings"
	"time"
)

func getWttr() string {
	wttrCmd := "curl -s 'wttr.in/?format=2'"
	wttrRun, err := exec.Command("sh", "-c", wttrCmd).Output()
	if err != nil {
		fmt.Println(err)
	}
	return strings.TrimSpace(string(wttrRun))
}

func getRam() string {
	ramCmd := "free -h | gawk '/Mem:/ {print $4}'"
	ramGib, err := exec.Command("sh", "-c", ramCmd).Output()
	if err != nil {
		fmt.Println(err)
	}
	return strings.TrimSpace(string(ramGib))
}

func main() {
	// initial run
	i := 0
	wttr := getWttr()
	ramFree := getRam()

	for i < 3700 {
		hTime := time.Now().Format("Jan 02, 2006 15:04")

		// get weather once per hour
		if i == 3600 {
			wttr = getWttr()
		}

		// get free Ram every 3 seconds
		if i%3 == 0 {
			ramFree = getRam()
		}

		// storing desired items as strings
		cat := []string{"RAM:", ramFree, "free", "|", wttr, "|", hTime}

		// concatinate all strings to one line for output
		status := strings.Join(cat, " ")

		cmd := exec.Command("xsetroot", "-name", status)
		cmd.Run()
		time.Sleep(1 * time.Second)

		// reset i to loop wttr updates
		if i > 3600 {
			i = 0
		} else {
			i++
		}
	}
}
