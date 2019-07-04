package icons

import (
	"fmt"
	"testing"
)

// These tests are for icon functions that rely on other data.
// Any icon that simply displays no matter the status is now tested.

func TestDropbox(t *testing.T) {
	db, err := Dropbox()
	if err != nil {
		t.Error(err) //.Error())
	}
	fmt.Println(db, err)
}

func TestRedshift(t *testing.T) {
	rs, err := Redshift()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(rs, err)
}

func TestVolume(t *testing.T) {
	vi, err := Volume()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(vi)
}

func TestSyncthing(t *testing.T) {
	sync, err := Syncthing()
	if err != nil {
		t.Error(err) //.Error())
	}
	fmt.Println(sync, err)
}
