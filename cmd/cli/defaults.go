package cli

import (
	"fmt"
	"path/filepath"

	"github.com/PederHA/d2herogrid/pkg/model"
	"github.com/mitchellh/go-homedir"
)

// UserConfig defaults
var (
	DefaultBrackets      = model.Brackets{model.BracketImmortal}
	DefaultGridName      = "d2hg"
	DefaultLayout        = model.LayoutMainStat
	DefaultSortAscending = false
	// Default Path is managed within NewUserConfigDefaults
)

var userconfigDir string

func init() {
	dir, err := homedir.Dir()
	if err != nil {
		// NOTE: Not panicing here
		fmt.Printf("Unable to detect home directory! A persistent config cannot be saved/loaded.")
	} else {
		userconfigDir = filepath.Join(dir, ".d2herogrid")
	}
}
