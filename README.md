# Go Status Bar
This is a script written with Go to display information. Currently tested with DWM and Tmux

## How To Use:
1. Open main.go in a text editor
2. Edit the last `status` variable to contain the blocks you wish to use
3. Change the laste line to match your UI (eg, `ui.Dwm(status)`)
4. Build the binary
5. Edit your config file to use the new binary

### Tmux Extras:
- Add to Tmux:
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
#### Blocks:
- Date & Time
- Weather via wttr.in
- Total RAM not used
#### User Interfaces:
- DWM
- I3WM
- Tmux
