package main

import (
	//"fmt"
	"os/exec"
	"time"
)

func main() {

	for true {
		hTime := time.Now().Format("2006-01-02 15:04:05")
		//fmt.Print(hTime.String(), "\n")
		cmd := exec.Command("xsetroot", "-name", hTime)
		cmd.Run()
		time.Sleep(1 * time.Second)
	}
}
