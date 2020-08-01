package model

// Layout represents a hero grid layout
type Layout struct {
	Name       string   `json:"layout" yaml:"layout"`
	Aliases    []string `json:"-" yaml:"-"` // CLI Aliases
	Categories []string `json:"-" yaml:"-"`
	// categoryIdx map[string]int 			// ?
	//gridFunc func(...) error 				// Should have something like this
}

var (
	// LayoutMainStat is the default Dota 2 hero grid layout, which is divided
	// into 3 categories: Strength, Agility & Intelligence.
	LayoutMainStat = &Layout{
		Name:       "Main Stat",
		Aliases:    []string{"mainstat", "ms", "stat"},
		Categories: []string{"Strength", "Agility", "Intelligence"},
	}
	// LayoutSingle is a layout where all heroes are gathered into a single category
	LayoutSingle = &Layout{
		Name:       "Single",
		Aliases:    []string{"single", "s"},
		Categories: []string{"Heroes"},
	}
	// LayoutAttackType is a layout where heroes are categorized based on attack type,
	// which is either Melee or Ranged.
	LayoutAttackType = &Layout{
		Name:       "Attack Type",
		Aliases:    []string{"attack", "attacktype", "attack type", "a"},
		Categories: []string{"Melee", "Ranged"},
	}
	// LayoutRole is a layout where heroes are categorized based on their typical
	// in-game role. Which is divided into the following 3 categories:
	// Carry, Support & Flex.
	LayoutRole = &Layout{
		Name:       "Role",
		Aliases:    []string{"role", "r"},
		Categories: []string{"Carry", "Support", "Flex"},
	}
	// LayoutLegs is a layout where heroes are categorized based on their number of legs.
	LayoutLegs = &Layout{
		Name:       "Legs",
		Aliases:    []string{"legs", "l"},
		Categories: []string{"0 Legs", "2 Legs", ">3 Legs"},
	}
	// LayoutModify is a special type of layout that signifies that d2herogrid should
	// attempt to sort the categories of an existing hero grid.
	LayoutModify = &Layout{
		Name:       "Modify",
		Aliases:    []string{"modify", "m", "none", "n"},
		Categories: []string{"none"},
	}
)

// AllLayouts is a slice of all layouts
var AllLayouts = []*Layout{
	LayoutMainStat,
	LayoutSingle,
	LayoutAttackType,
	LayoutRole,
	LayoutLegs,
	LayoutModify,
}
