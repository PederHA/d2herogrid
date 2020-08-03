// +build windows

package cli

import (
	"testing"
)

func TestGetSteamPathWindows(t *testing.T) {
	p, err := getSteamPathWindows()
	if err != nil {
		t.Error(err)
	}
	if p == "" {
		t.Errorf("getSteamPathWindows() returned empty string and no error.")
	}
}
