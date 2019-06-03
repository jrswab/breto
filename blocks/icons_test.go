package blocks

import (
	"fmt"
	"testing"
)

func TestDropboxIcon(t *testing.T) {
	db, err := DropboxIcon()
	if err != nil {
		t.Error(err) //.Error())
	}
	fmt.Println(db, err)
}

func TestRedshiftIcon(t *testing.T) {
	rs, err := RedshiftIcon()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(rs, err)
}

func TestVolumeIcon(t *testing.T) {
	vi, err := VolumeIcon()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(vi)
}
