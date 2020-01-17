package icons

import (
	"fmt"
	"testing"
)

func TestEncodeEmoji(t *testing.T) {
	var emojis = []struct {
		input    string
		expected string
	}{
		{"00001F4E5", "📥 "},
		{"00001F4A1", "💡 "},
		{"000002194", "↔ "},
	}

	for _, test := range emojis {
		output := encodeEmoji(test.input)
		if output != test.expected {
			errMessage := fmt.Sprintf("Expected %s, got %s", test.expected, output)
			t.Error(errMessage)
		}
	}
}

func TestDropbox(t *testing.T) {
	db, err := Dropbox(false)
	if err != nil {
		t.Error(err) //.Error())
	}
	fmt.Println(db, err)
}

func TestRedshift(t *testing.T) {
	rs, err := Redshift(false)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(rs, err)
}

func TestVolume(t *testing.T) {
	vi, err := Volume(false)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(vi)
}

func TestSyncthing(t *testing.T) {
	sync, err := Syncthing(false)
	if err != nil {
		t.Error(err) //.Error())
	}
	fmt.Println(sync, err)
}
