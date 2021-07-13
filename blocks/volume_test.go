package blocks

import (
	"fmt"
	"testing"
)

func TestVolumeText(t *testing.T) {
	vol, err := Volume(false)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(vol)
}
