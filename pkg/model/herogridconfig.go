package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"sort"
	"strings"

	"github.com/PederHA/d2herogrid/internal/utils"
	"github.com/PederHA/d2herogrid/pkg/config"
)

// HeroGridConfig represents the contents of hero_grid_config.json
type HeroGridConfig struct {
	config.Config
	Version   int        `json:"version"`
	HeroGrids []HeroGrid `json:"configs"`
}

// NewHeroGridConfig creates a new HeroGridConfig by reading the contents
// of an existing hero_grid_config.json file.
func NewHeroGridConfig(hgcPath string) (*HeroGridConfig, error) {
	var hgc = &HeroGridConfig{}
	f, err := os.Open(hgcPath)
	defer f.Close()
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, hgc)
	if err != nil {
		return nil, err
	}

	return hgc, nil
}

// NewHeroGridConfigDefault returns a new HeroGridConfig using default Dota 2 parameters,
// which is simply an empty config.
func NewHeroGridConfigDefault() *HeroGridConfig {
	return &HeroGridConfig{
		Version:   3,
		HeroGrids: []HeroGrid{},
	}
}

// SaveConfigJSON saves the contents of a HeroGridConfig as a JSON-encoded file
func (h *HeroGridConfig) SaveConfigJSON(path string) error {
	return utils.MarshalJSON(path, h)
}

// ListGrids lists the names of all grids and their categories in a HeroGridConfig
// NOTE: Unused!
func (h *HeroGridConfig) ListGrids() {
	fmt.Printf("Grids:\n")
	for _, config := range h.HeroGrids {
		fmt.Printf("\tConfig Name: %s\n", config.ConfigName)
		for _, category := range config.Categories {
			fmt.Printf("\t\tCategory Name: %s\n", category.CategoryName)
		}
	}
}

func (h *HeroGridConfig) String() string {
	var s string
	s += fmt.Sprintf("Version: %d\n", h.Version)
	for _, herogrid := range h.HeroGrids {
		s += herogrid.String()
	}
	return s
}

// MakeGrid creates a new hero grid or modifies an existing one.
// After processing the grid, it adds the grid to HeroGridConfig's list of grids.
func (h *HeroGridConfig) MakeGrid(gridName string, layout *Layout, bracket *Bracket, heroes *Heroes) error {
	var gridFunc func(*HeroGrid, *Heroes) (*HeroGrid, error)

	switch layout {
	case LayoutSingle:
		gridFunc = h.newSingleGrid
	case LayoutMainStat:
		gridFunc = h.newMainStatGrid
	case LayoutAttackType:
		gridFunc = h.newAttackTypeGrid
	case LayoutRole:
		gridFunc = h.newRoleGrid
	case LayoutLegs:
		gridFunc = h.newLegsGrid
	case LayoutModify:
		gridFunc = h.modifyGrid
	default:
		return fmt.Errorf("model: encountered unknown layout '%s' when attempting to generate hero grid. This should NEVER happen", layout)
	}

	var cfgName string
	if layout == LayoutModify {
		cfgName = gridName
	} else {
		cfgName = getCfgName(gridName, bracket.Name, layout.Name)
	}

	// Make new grid
	grid, err := NewHeroGrid(cfgName, layout.Categories)
	if err != nil {
		return err
	}

	// Populate grid
	grid, err = gridFunc(grid, heroes)
	if err != nil {
		return err
	}

	// Add grid to HeroGridConfig
	err = h.addHeroGrid(grid)
	if err != nil {
		return err
	}
	return nil
}

func (h *HeroGridConfig) findGridIdx(gridName string) (int, bool) {
	for idx, cfg := range h.HeroGrids {
		if cfg.ConfigName == gridName {
			return idx, true
		}
	}
	return -1, false
}

// UNUSED
func (h *HeroGridConfig) getHeroGridByName(gridName string) (*HeroGrid, error) {
	if idx, ok := h.findGridIdx(gridName); ok {
		herogrid := h.HeroGrids[idx]
		return &herogrid, nil
	}
	return nil, fmt.Errorf("unable to find a hero grid with the name '%s'", gridName)
}

////////////////////////////////////////////////////////////////////////////////
// It doesn't REALLY make sense that these are HeroGridConfig methods
////////////////////////////////////////////////////////////////////////////////

func (h *HeroGridConfig) newSingleGrid(grid *HeroGrid, heroes *Heroes) (*HeroGrid, error) {
	for _, hero := range *heroes {
		grid.Categories[0].HeroIDs = append(grid.Categories[0].HeroIDs, hero.HeroID)
	}
	return grid, nil
}

func (h *HeroGridConfig) newMainStatGrid(grid *HeroGrid, heroes *Heroes) (*HeroGrid, error) {
	var categoryIdx = map[string]int{"str": 0, "agi": 1, "int": 2}
	for _, hero := range *heroes {
		idx := categoryIdx[hero.PrimaryAttr] // NOTE: no error checking here
		grid.Categories[idx].HeroIDs = append(grid.Categories[idx].HeroIDs, hero.HeroID)
	}
	return grid, nil
}

func (h *HeroGridConfig) newAttackTypeGrid(grid *HeroGrid, heroes *Heroes) (*HeroGrid, error) {
	var categoryIdx = map[string]int{"melee": 0, "ranged": 1}
	for _, hero := range *heroes {
		idx := categoryIdx[hero.AttackType]
		grid.Categories[idx].HeroIDs = append(grid.Categories[idx].HeroIDs, hero.HeroID)
	}
	return grid, nil
}

func (h *HeroGridConfig) newRoleGrid(grid *HeroGrid, heroes *Heroes) (*HeroGrid, error) {
	return nil, errors.New("Role grid is not yet implemented")
}

func (h *HeroGridConfig) newLegsGrid(grid *HeroGrid, heroes *Heroes) (*HeroGrid, error) {
	return nil, errors.New("Legs grid is not yet implemented")
}

// Below is the (shoddy) implementation of modifying an existing grid in place by
// re-arranging heroes in each category based on their winrate

type hero struct {
	heroID int
	index  int
}

type heroesGrid []hero

func (h heroesGrid) Len() int           { return len(h) }
func (h heroesGrid) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h heroesGrid) Less(i, j int) bool { return h[i].index < h[j].index }

func (h *HeroGridConfig) modifyGrid(grid *HeroGrid, heroes *Heroes) (*HeroGrid, error) {
	// Find a grid with the specified name
	grid, err := h.getHeroGridByName(grid.ConfigName)
	if err != nil {
		return nil, err
	}

	// Make mapping of each heroID and their respective absolute ranking
	heroesIdx := make(map[int]int) // map[heroID]rankingIdx
	for idx, hero := range *heroes {
		heroesIdx[hero.HeroID] = idx
	}
	var categories []Category // slice of categories to replace existing categories with

	// Modify the existing grid
	for _, cat := range grid.Categories {
		var heroesInCategory heroesGrid // Every hero in the category has an ID and an index
		for _, heroID := range cat.HeroIDs {
			if idx, ok := heroesIdx[heroID]; ok {
				heroesInCategory = append(heroesInCategory, hero{heroID, idx})
			} else {
				return nil, fmt.Errorf(
					"encountered unknown hero ID %d when attemptng to modify the grid %s",
					idx, grid.ConfigName,
				)
			}
		}

		// Sort heroes in category based their index (winrate)
		sort.Sort(heroesInCategory)
		var heroIDs []int
		for _, h := range heroesInCategory {
			heroIDs = append(heroIDs, h.heroID)
		}
		cat.HeroIDs = heroIDs
		categories = append(categories, cat)
	}
	grid.Categories = categories // replace unsorted categories with new sorted categories
	return grid, nil
}

func (h *HeroGridConfig) modifyGridOriginal(grid *HeroGrid, heroes *Heroes) (*HeroGrid, error) {
	// Find a grid with the specified name
	grid, err := h.getHeroGridByName(grid.ConfigName)
	if err != nil {
		return nil, err
	}

	var categories []Category
	// Modify the existing grid
	for _, cat := range grid.Categories {
		var heroesCat heroesGrid // Every hero in the category has an ID and an index
		for _, heroID := range cat.HeroIDs {
			for idx, h := range *heroes { // param heroes is sorted by winrate
				if heroID == h.HeroID {
					heroesCat = append(heroesCat, hero{h.HeroID, idx})
					break
				}
			}
		}
		// Sort heroes in category based their index (winrate)
		sort.Sort(heroesCat)
		var heroIDs []int
		for _, h := range heroesCat {
			heroIDs = append(heroIDs, h.heroID)
		}
		cat.HeroIDs = heroIDs
		categories = append(categories, cat)
	}
	grid.Categories = categories
	return grid, nil
}

func (h *HeroGridConfig) addHeroGrid(grid *HeroGrid) error {
	if gridIdx, ok := h.findGridIdx(grid.ConfigName); ok {
		h.HeroGrids[gridIdx] = *grid
	} else {
		h.HeroGrids = append(h.HeroGrids, *grid)
	}
	return nil
}

// HeroGrid represents a single hero grid, which contains optionally contains
// 1 or more category.
type HeroGrid struct {
	config.Config
	ConfigName string     `json:"config_name"`
	Categories []Category `json:"categories"`
}

// NewHeroGrid creates a new HeroGrid from a name and a list of categories
func NewHeroGrid(name string, categories []string) (*HeroGrid, error) {
	var h = &HeroGrid{ConfigName: name}

	var cat []Category
	switch len(categories) {
	case 1:
		cat = singleCategory
	case 2:
		cat = doubleCategory
	case 3:
		cat = tripleCategory
	default:
		return nil, fmt.Errorf("model.NewHeroGrid: expected <=3 categories, received %d", len(categories))
	}

	h.Categories = make([]Category, len(categories))
	copy(h.Categories, cat)

	// Assert that length of categories argument matches length of categories from switch..case
	if len(categories) != len(h.Categories) {
		panic("model.NewHeroGrid: Mismatched length of categories. This should never happen.")
	}

	for i, n := range categories {
		h.Categories[i].CategoryName = n
	}
	return h, nil
}

func (h *HeroGrid) String() string {
	var s string
	s += fmt.Sprintf("ConfigName: %s\n", h.ConfigName)
	for _, category := range h.Categories {
		s += category.String()
	}
	return s
}

// Category represents a single category within a hero grid
type Category struct {
	config.Config
	CategoryName string  `json:"category_name"`
	XPosition    float64 `json:"x_position"`
	YPosition    float64 `json:"y_position"`
	Width        float64 `json:"width"`
	Height       float64 `json:"height"`
	HeroIDs      []int   `json:"hero_ids"` // uint8?
}

// NewCategory creates a new Category
func NewCategory(name string, xpos, ypos, width, height float64) Category {
	return Category{
		CategoryName: name,
		XPosition:    math.Abs(xpos),
		YPosition:    math.Abs(ypos),
		Width:        math.Abs(width),
		Height:       math.Abs(height),
	}
}

// NewCategoryDefault creates a new Category with default parameters
func NewCategoryDefault(name string, xpos, ypos, width, height float64) Category {
	return Category{
		CategoryName: "name",
		XPosition:    0.0,
		YPosition:    0.0,
		Width:        1180.0,
		Height:       180.0,
	}
}

func (c *Category) String() string {
	var b *strings.Builder
	b.Grow(256)
	fmt.Fprintf(b, "CategoryName: %s\n", c.CategoryName)
	fmt.Fprintf(b, "XPosition: %f\n", c.XPosition)
	fmt.Fprintf(b, "YPosition: %f\n", c.YPosition)
	fmt.Fprintf(b, "Width: %f\n", c.Width)
	fmt.Fprintf(b, "Height: %f\n", c.Height)
	if len(c.HeroIDs) > 0 {
		fmt.Fprint(b, "[")
		for i, id := range c.HeroIDs {
			fmt.Fprintf(b, "%d", id)
			if i != (len(c.HeroIDs) - 1) {
				fmt.Fprint(b, ", ")
			}
		}
		fmt.Fprintf(b, "]")

	}
	return b.String()
}

func getCfgName(basename, bracket, layout string) string {
	return fmt.Sprintf("%s (%s) (%s)", basename, bracket, layout)
}

var (
	singleCategory = []Category{
		NewCategory("One", 0, 0, 1180, 1180),
	}

	doubleCategory = []Category{
		NewCategory("One", 0, 0, 1180, 280),
		NewCategory("Two", 0, 300, 1180, 280),
	}

	tripleCategory = []Category{
		NewCategory("One", 0, 0, 1180, 180),
		NewCategory("Two", 0, 200, 1180, 180),
		NewCategory("Three", 0, 400, 1180, 180),
	}
)
