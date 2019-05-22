# DWM Status Bar
This is a script written with Go to display information for DWM via `xsetroot -name`.

The reason for using Go is because when a using Posix compliant shell script, 
the execition of programs slowed down noticably.
This may not be the script itself and something else but after moving to a 
compiled binary written in go the problem halted.

## To run:
1. Download the source
1. Edit as desired
1. `go build`
2. Execute the new binary

### Optional:
1. Set DWM to run the new binary automatically.
2. Set keycombo to run the binary.

## Current Features:
### Master Branch:
- Date & Time