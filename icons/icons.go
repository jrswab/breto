package icons

import (
	"encoding/binary"
	"fmt"
	"os/exec"
)

// Dropbox sends the Dropbox icon when the app is running.
func Dropbox(emoji bool) (string, error) {
	// ps ux | awk '/dropbox/ {print $11}' | grep dropbox
	dbIcon := ""
	if emoji {
		dbIcon = "🗄"
	}
	dbCmd := "ps ux | gawk '/dropbox/ {print $11}' | grep dropbox"
	runDbCmd, err := exec.Command("sh", "-c", dbCmd).Output()

	if err != nil && err.Error() != "exit status 1" {
		return "", err
	}
	if string(runDbCmd) != "" {
		return fmt.Sprintf("%s", dbIcon), nil
	}
	return "", nil
}

// Redshift sends the icon when the app is running.
func Redshift(emoji bool) (string, error) {
	rsIcon := ""
	if emoji {
		rsIcon = "💡"
	}
	rsCmd := "ps ux | gawk '/redshift/ {print $11}' | grep redshift"
	runRsCmd, err := exec.Command("sh", "-c", rsCmd).Output()

	if err != nil && err.Error() != "exit status 1" {
		return "", err
	}
	if string(runRsCmd) != "" {
		return fmt.Sprintf("%s", rsIcon), nil
	}
	return "", nil
}

// Syncthing sends the icon when the app is running.
func Syncthing(emoji bool) (string, error) {
	syncIcon := ""
	if emoji {
		syncIcon = "⬌"
	}
	syncCmd := "ps ux | gawk '/syncthing/ {print $11}' | grep syncthing"
	runSyncCmd, err := exec.Command("sh", "-c", syncCmd).Output()

	if err != nil && err.Error() != "exit status 1" {
		return "", err
	}
	if string(runSyncCmd) != "" {
		return fmt.Sprintf("%s", syncIcon), nil
	}
	return "", nil
}

// Volume sends the icon when the app is running.
func Volume(emoji bool) (string, error) {
	// Font Awesome:
	volIconMute := " "
	volIconLow := " "
	volIconMid := " "
	volIconHigh := " "

	// Openmoji:
	if emoji {
		volIconMute = "🔇 "
		volIconLow = "🔈 "
		volIconMid = "🔉 "
		volIconHigh = "🔊 "
	}

	volCmd := "amixer -D pulse sget Master | awk '/Front Right:/ {print $5}' | grep -o '[0-9]*'"
	runVolCmd, err := exec.Command("sh", "-c", volCmd).Output()
	volValue := binary.LittleEndian.Uint16(runVolCmd)

	if err != nil && err.Error() != "exit status 1" {
		return "", err
	}

	switch {
	case volValue == 0:
		return fmt.Sprintf("%s", volIconMute), nil
	case volValue < 50:
		return fmt.Sprintf("%s", volIconLow), nil
	case volValue >= 50 && volValue <= 74:
		return fmt.Sprintf("%s", volIconMid), nil
	case volValue >= 75:
		return fmt.Sprintf("%s", volIconHigh), nil
	default:
		return "", nil
	}
}

// The following have no checks that need to be made

// Dir sends the icon when the app is running.
func Dir(emoji bool) string {
	if emoji {
		return "📂 "
	}
	return " "
}

// Mem sends the icon when the app is running.
func Mem(emoji bool) string {
	if emoji {
		return "💾 "
	}
	return " "
}

// Temp sends the icon when the app is running.
func Temp(emoji bool) string {
	if emoji {
		return "🌡 "
	}
	return " "
}

// Power sends the icon when the app is running.
func Power(emoji bool) string {
	if emoji {
		return "⚡ "
	}
	return " "
}
