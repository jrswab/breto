package ui

import (
	"fmt"
	"os/exec"
)

// Dwm - To have the status blocks appear in dwm we have to run xsetroot -name [status]
// where [status] contains the data blocks
func Dwm(status string) error {
	cmd := exec.Command("xsetroot", "-name", status)
	err := cmd.Run()
	return err
}

// Tmux - we only need to output the block to stdout
// Add the following to your tmux config:
// set -g status-right "#($HOME/PATHTO/go-status)"
// where `go-status` is the compiled binary
func Tmux(status string) {
	fmt.Println(status)
}

// I3wm - we only need to output the block to stdout as we do with tmux.
// Make sure to update the `bar {}` section in your i3wm
// config file to the following:
// status_command $HOME/PATH/TO/go-status
// where `go-status` is the compiled binary
func I3wm(status string) {
	fmt.Println(status)
}

// Polybar - we only need to output the block to stdout as we do with tmux.
// Add Font Awesome fonts to the Polybar config file.
func Polybar(status string) {
	fmt.Println(status)
}
