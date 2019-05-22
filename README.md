# DWM Status Bar
This is a script written with Go to display information for DWM via `xsetroot -name`.

The reason for using Go is because when a using Posix compliant shell script the execition of programs slowed down.
This may not be the script itself and something else but since moving to a compiled binary written in go the problem halted.

## To run:
1. Download the source
1. Edit as desired
1. `go build`
1. Set DWM to run the new binary.

## Current Features:
- Date & Time