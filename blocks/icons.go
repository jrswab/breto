package blocks

import (
	"encoding/binary"
	"fmt"
	"os/exec"
)

func DropboxIcon() (string, error) {
	// ps ux | awk '/dropbox/ {print $11}' | grep dropbox
	dbIcon := ""
	dbCmd := "ps ux | gawk '/dropbox/ {print $11}' | grep dropbox"
	runDbCmd, err := exec.Command("sh", "-c", dbCmd).Output()

	if err != nil && err.Error() != "exit status 1" {
		return "", err
	}
	if string(runDbCmd) != "" {
		return fmt.Sprintf("%s ", dbIcon), nil
	}
	return "", nil
}

func RedshiftIcon() (string, error) {
	rsIcon := ""
	rsCmd := "ps ux | gawk '/redshift/ {print $11}' | grep redshift"
	runRsCmd, err := exec.Command("sh", "-c", rsCmd).Output()

	if err != nil && err.Error() != "exit status 1" {
		return "", err
	}
	if string(runRsCmd) != "" {
		return fmt.Sprintf("%s ", rsIcon), nil
	}
	return "", nil
}

func VolumeIcon() (string, error) {
	volIconMute := ""
	volIconLow := ""
	volIconMid := ""
	volIconHigh := ""
	volCmd := "amixer -D pulse sget Master | awk '/Front Right:/ {print $5}' | grep -o '[0-9]*'"
	runVolCmd, err := exec.Command("sh", "-c", volCmd).Output()
	volValue := binary.LittleEndian.Uint16(runVolCmd)

	if err != nil && err.Error() != "exit status 1" {
		return "", err
	}

	switch {
	case volValue == 0:
		return fmt.Sprintf("%s ", volIconMute), nil
	case volValue < 50:
		return fmt.Sprintf("%s ", volIconLow), nil
	case volValue >= 50 && volValue <= 74:
		return fmt.Sprintf("%s ", volIconMid), nil
	case volValue >= 75:
		return fmt.Sprintf("%s ", volIconHigh), nil
	default:
		return "", nil
	}
}
