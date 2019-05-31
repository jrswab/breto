# Go Status Bar
This is a script written with Go to display information. Currently tested with DWM and Tmux

## Running in DWM:
1. Download the source
1. Edit as desired
1. `go build`
2. Execute the new binary

### DWM Optionals:
1. Set DWM to run the new binary automatically.
2. Set keycombo to run the binary.

## Running in Tmux:
1. Comment out or remove:
   - `os/exec`
   - `var cmd *exec.Cmd`
   - `cmd = exec.Command("xsetroot", "-name", status)`
   - `cmd.Run()`
2. Add `fmt.Println(status)` to the end of `for range ticker.C`
3. Save and run `go build` within the project directory
4. Add to Tmux:
   - `set -g status-right "#($HOME/PATHTO/tmux-status)"`
	   - if you have colors in this setting add the path at the end of the string
	   - Be sure to use the correct path and name of the file you built with GO.
	   - Running `mv dwm-status ~/tmux-status` will allow you to use `"($HOME/tmux-status)" in your config.
   - `set -g status-right-length 53`
     - if you are not using all the custom packages this number can be lower
	 - 53 worked well for my setup but if you notice the status getting cut off just increase the number and reload tmux.


## Wttr.in Optionals:
1. Add your area to the weather function
   - Area Code: 'wttr.in/~00000?format=2' 
   - City: 'wttr.in/~Paris?format=2'

## Current Features:
### Master Branch:
- Date & Time
- Weather via wttr.in
- Total RAM not used
