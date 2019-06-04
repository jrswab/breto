# Go Status Bar
This is a script written with Go to display information.
Currently tested with DWM, i3wm, and tmux.

## Current Features:
### Master Branch:
#### Blocks:
- Date & Time
- Weather via wttr.in
- Total RAM not used
- Icons
#### Current UIs:
- DWM
- i3wm
- tmux

### Icons:
To display icons in DWM.
1. Install FontAwesome v4
2. Set FontAwesome as the second font in `dwm/config.h`
   - eg. `static const char *fonts[] = { "Source Code Pro:size=13", "FontAwesome:size=14" };`
	 - Relaunch DWM

Currently, the icons are much too small to be usefull in Tmux or i3wm without further configuration.
If you have an easy way to display FontAwesome icons at the same scale as the terminal text please submit a pull request.

## How To Use:
1. Open main.go in a text editor.
1. If cloned from Github change the custom package directories from Gitlab to Github.
2. Edit the last `status` variable to contain the blocks you wish to use.
3. Change the last line to match your UI (eg, `ui.Dwm(status)`).
4. Build the binary.
5. Edit your config file to use the new binary.

### Adding to DWM:
- Simply set `ui.Dwm(status)`, compile, and execute.
- If you already have a startup script for DWM just add a new line for go-status

### Addng to Tmux:
- `set -g status-right "#($HOME/PATHTO/tmux-status)"`
 - If you have colors in this setting add the path at the end of the string
 - Be sure to use the correct path and name of the file you built with GO.
 - Running `mv go-status ~/tmux-status` will allow you to use `"($HOME/tmux-status)"` in your config.
- `set -g status-right-length 53`
 - If you are not using all the custom packages this number can be lower
 - This also will vary based on screen size. 53 is the minimum that worke for all current blocks
 - If you notice the status getting cut off just increase the number and reload tmux.

### Adding to i3wm:
- In the i3wm confige file, change `status_command ...` to `status_command PATH/TO/go-status`

## Wttr.in Options:
1. Add your area to the weather function
   - Area Code: 'wttr.in/~00000?format=2' 
   - City: 'wttr.in/~Paris?format=2'
	 - More information can be found at [wttr.in/:help](https://wttr.in/:help)
2. Add tweaks to `blocks/wttr.go`

## To-Do:
1. Get a proper name
2. Scale icons in tmux and i3wm
3. Add /home GiB free `df -h | awk '/home/ {print $4}`
