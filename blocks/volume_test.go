package blocks

import (
	"fmt"
	"testing"
)

func TestVolumeText(t *testing.T) {
	vol, err := VolumeText()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(vol)
}
