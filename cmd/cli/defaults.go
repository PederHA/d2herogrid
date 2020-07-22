// TODO: Rename this file
package cli

const (
	// Single combines all heroes into a single category
	LayoutSingle = iota
	// MainStat divides heroes into 3 categories Str/Agi/Int (default)
	LayoutMainStat
	// AttackType divides heroes into 2 categories of Melee/Ranged
	LayoutAttackType
	// Role divides heroes into 3 categories of Carry/Support/Flex
	LayoutRole
)

// UserConfig defaults
var defaultBrackets = []int{7, 8}
var defaultGridName = "OpenDota Hero Winrates"
var defaultLayout = LayoutMainStat
var defaultSortAscending = false

// Default Path is managed within NewUserConfigDefaults
