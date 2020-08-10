package cli

import (
	"github.com/PederHA/d2herogrid/pkg/model"
)

// UserConfig defaults
var (
	DefaultBrackets      = model.Brackets{model.BracketDivine}
	DefaultGridName      = "d2hg"
	DefaultLayout        = model.LayoutMainStat
	DefaultSortAscending = false
	// Default Path is managed within NewUserConfigDefaults
)
