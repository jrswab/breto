package main

import (
	"fmt"
	"os/exec"
	"strings"
	"time"
)

func main() {
	i := 0

	for i < 3700 {
		hTime := time.Now().Format("Jan 02, 2006 15:04")

		// get free Ram using Posix compliant shell command
		ramCmd := "free -h | gawk '/Mem:/ {print $4}'"
		ramGib, err := exec.Command("sh", "-c", ramCmd).Output()
		if err != nil {
			fmt.Println(err)
		}
		ramFree := strings.TrimSpace(string(ramGib))

		// storing desired items as strings
		cat := []string{"RAM:", ramFree,
			"free", "|", hTime}

		// concatinate all strings to one line for output
		status := strings.Join(cat, " ")

		cmd := exec.Command("xsetroot", "-name", status)
		cmd.Run()
		time.Sleep(1 * time.Second)

		if i > 3600 {
			i = 0
		} else {
			i++
		}
	}
}
