package model

// Bracket represents a Dota skill bracket
type Bracket struct {
	Name       string   `json:"bracket" yaml:"bracket"`
	Aliases    []string `json:"-" yaml:"-"` // CLI Aliases
	Categories []string `json:"-" yaml:"-"`
	// categoryIdx map[string]int 			// ?
	//gridFunc func(...) error 				// Should have something like this
}

// Brackets is a slice of Bracket pointers
type Brackets []*Bracket

var (
	// BracketHerald represents the Herald skill bracket
	BracketHerald = &Bracket{
		Name:    "Herald",
		Aliases: []string{"herald", "h", "1"},
	}
	// BracketGuardian represents the Guardian skill bracket
	BracketGuardian = &Bracket{
		Name:    "Guardian",
		Aliases: []string{"guardian", "g", "2"},
	}
	// BracketCrusader represents the Crusader skill bracket
	BracketCrusader = &Bracket{
		Name:    "Crusader",
		Aliases: []string{"crusader", "c", "3"},
	}
	// BracketArchon represents the Archon skill bracket
	BracketArchon = &Bracket{
		Name:    "Archon",
		Aliases: []string{"archon", "a", "4"},
	}
	// BracketLegend represents the Legend skill bracket
	BracketLegend = &Bracket{
		Name:    "Legend",
		Aliases: []string{"legend", "l", "5"},
	}
	// BracketAncient represents the Ancient skill bracket
	BracketAncient = &Bracket{
		Name:    "Ancient",
		Aliases: []string{"ancient", "A", "6"},
	}
	// BracketDivine represents the Divine skill bracket
	BracketDivine = &Bracket{
		Name:    "Divine",
		Aliases: []string{"divine", "d", "7"},
	}
	// BracketImmortal represents the Immortal skill bracket
	BracketImmortal = &Bracket{
		Name:    "Immortal",
		Aliases: []string{"immortal", "i", "8"},
	}
	// BracketPro represents official pro matches
	BracketPro = &Bracket{
		Name:    "Pro",
		Aliases: []string{"pro", "p", "9"},
	}
)

// AllBrackets is a slice of all brackets
var AllBrackets = Brackets{
	BracketHerald,
	BracketGuardian,
	BracketCrusader,
	BracketArchon,
	BracketLegend,
	BracketAncient,
	BracketDivine,
	BracketImmortal,
	BracketPro,
}
