package ram

import (
	"fmt"
	"os/exec"
	"strings"
	"time"
)

func Free(cRam chan string) {
	ramFree := ""
	for i := 0; i < 5; {
		if i == 0 || i == 3 {
			// set shell command
			ramCmd := "free -h | gawk '/Mem:/ {print $4}'"
			// run and save the output of ramCmd to ramGib
			ramGib, err := exec.Command("sh", "-c", ramCmd).Output()
			if err != nil {
				cRam <- err.Error()
			}
			// return ramGib as string
			ramFree = fmt.Sprintf("Ram: %s free | ", strings.TrimSpace(string(ramGib)))
		}
		cRam <- ramFree
		if i == 3 {
			i = 0
		} else {
			i++
		}
		time.Sleep(1 * time.Second)
	}
}
