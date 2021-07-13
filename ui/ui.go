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

// Default is used when we only need to output the block to stdout.
// This is the default option.
//
// For tmux:
// Add the following to your tmux config:
// set -g status-right "#($HOME/PATH/TO/breto)"
// where `breto` is the compiled binary
//
// For i3wm:
// Make sure to update the `bar {}` section in your i3wm
// config file to the following:
// status_command $HOME/PATH/TO/breto
// where `breto` is the compiled binary
//
// For Polybar:
// [module/breto]
// type = custom/script
// exec = /path/to/breto/binary
// tail = true
func Default(status string) {
	fmt.Println(status)
}
