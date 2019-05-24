# DWM Status Bar
This is a script written with Go to display information for DWM via `xsetroot -name`.

The reason for using Go is because when a using Posix compliant shell script, the execution of programs slowed down noticeably.
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
3. Add your area to the weather function
   - Area Code: 'wttr.in/~00000?format=2' 
   - City: 'wttr.in/~Paris?format=2'

## Current Features:
### Master Branch:
- Date & Time
- Weather via wttr.in
- Total RAM not used