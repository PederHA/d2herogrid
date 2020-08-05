package model

// Layout represents a hero grid layout
type Layout struct {
	Name       string     `json:"layout" yaml:"layout"`
	Aliases    []string   `json:"-" yaml:"-"` // CLI Aliases
	Categories []Category `json:"-" yaml:"-"` // TODO: Should be []Category
	// categoryIdx map[string]int 			// ?
	//gridFunc func(...) error 				// Should have something like this
}

var (
	// LayoutMainStat is the default Dota 2 hero grid layout, which is divided
	// into 3 categories: Strength, Agility & Intelligence.
	LayoutMainStat = &Layout{
		Name:    "Main Stat",
		Aliases: []string{"mainstat", "ms", "stat"},
		Categories: []Category{
			NewCategory("Strength", 0, 0, 1180, 180),
			NewCategory("Agility", 0, 200, 1180, 180),
			NewCategory("Intelligence", 0, 400, 1180, 180),
		},
	}
	// LayoutSingle is a layout where all heroes are gathered into a single category
	LayoutSingle = &Layout{
		Name:    "Single",
		Aliases: []string{"single", "s"},
		Categories: []Category{
			NewCategory("Heroes", 0, 0, 1180, 1180),
		},
	}
	// LayoutAttackType is a layout where heroes are categorized based on attack type,
	// which is either Melee or Ranged.
	LayoutAttackType = &Layout{
		Name:    "Attack Type",
		Aliases: []string{"attack", "attacktype", "attack type", "a"},
		Categories: []Category{
			NewCategory("Melee", 0, 0, 1180, 280),
			NewCategory("Ranged", 0, 300, 1180, 280),
		},
	}
	// LayoutRole is a layout where heroes are categorized based on their typical
	// in-game role. Which is divided into the following 3 categories:
	// Carry, Support & Flex.
	LayoutRole = &Layout{
		Name:    "Role",
		Aliases: []string{"role", "r"},
		Categories: []Category{
			NewCategory("Carry", 0, 0, 1180, 240),
			NewCategory("Support", 0, 260, 1180, 180),
			NewCategory("Flexible", 0, 460, 1180, 95),
		},
	}
	// LayoutLegs is a layout where heroes are categorized based on their number of legs.
	LayoutLegs = &Layout{
		Name:    "Legs",
		Aliases: []string{"legs", "l"},
		Categories: []Category{
			NewCategory("0 Legs", 0, 0, 1180, 95),
			NewCategory("2 Legs", 0, 130, 1180, 315),
			NewCategory(">= 4 Legs", 0, 480, 1180, 110),
		},
	}
	// LayoutModify is a special type of layout that signifies that d2herogrid should
	// attempt to sort the categories of an existing hero grid.
	LayoutModify = &Layout{
		Name:    "Modify",
		Aliases: []string{"modify", "m", "none", "n"},
		Categories: []Category{
			NewCategory("Heroes", 0, 0, 0, 0),
		},
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
