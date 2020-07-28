// TODO: Rename this file
package cli

import (
	"fmt"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
)

const (
	// LayoutSingle combines all heroes into a single category
	LayoutSingle = "single"
	// LayoutMainStat divides heroes into 3 categories Str/Agi/Int (default)
	LayoutMainStat = "mainstat"
	// LayoutAttackType divides heroes into 2 categories of Melee/Ranged
	LayoutAttackType = "mttackype"
	// LayoutRole divides heroes into 3 categories of Carry/Support/Flex
	LayoutRole = "role"
	// LayoutLegs divides heroes into categories based on number of legs
	LayoutLegs = "legs"
	// LayoutNone modifies an existing hero grid without changing the layout
	LayoutNone = "none"
)

var Layouts = map[string]bool{
	LayoutSingle:     true,
	LayoutMainStat:   true,
	LayoutAttackType: true,
	LayoutRole:       true,
	LayoutLegs:       true,
	LayoutNone:       true,
}

const (
	BracketHerald   = "Herald"
	BracketGuardian = "Guardian"
	BracketCrusader = "Crusader"
	BracketArchon   = "Archon"
	BracketLegend   = "Legend"
	BracketAncient  = "Ancient"
	BracketDivine   = "Divine"
	BracketImmortal = "Immortal"
	BracketPro      = "Pro"
)

var Brackets = map[string]bool{
	BracketHerald:   true,
	BracketGuardian: true,
	BracketCrusader: true,
	BracketArchon:   true,
	BracketLegend:   true,
	BracketAncient:  true,
	BracketDivine:   true,
	BracketImmortal: true,
	BracketPro:      true,
}

// UserConfig defaults
var (
	defaultBrackets      = []string{"Divine", "Immortal"}
	defaultGridName      = "d2hg"
	defaultLayout        = LayoutMainStat
	defaultSortAscending = false
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
