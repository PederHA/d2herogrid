package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"os"

	"github.com/PederHA/d2herogrid/cmd/cli"
	"github.com/PederHA/d2herogrid/pkg/config"
)

// HeroGridConfig represents the contents of hero_grid_config.json
type HeroGridConfig struct {
	config.Config
	Version   int        `json:"version"`
	HeroGrids []HeroGrid `json:"configs"`
}

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

func (h *HeroGridConfig) SaveConfigJSON(path string) error {
	b, err := json.MarshalIndent(h, "", "    ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path, b, 0644)
	if err != nil {
		return err
	}

	return nil
}

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

func (h *HeroGridConfig) MakeGrid(gridName string, layout string, bracket string, heroes *Heroes) error {
	var gridFunc func(string, string, *Heroes) (*HeroGrid, error)
	switch layout {
	case cli.LayoutSingle:
		gridFunc = h.newSingleGrid
	case cli.LayoutMainStat:
		gridFunc = h.newMainStatGrid
	case cli.LayoutAttackType:
		gridFunc = h.newAttackTypeGrid
	case cli.LayoutRole:
		gridFunc = h.newRoleGrid
	case cli.LayoutLegs:
		gridFunc = h.newLegsGrid
	case cli.LayoutNone:
		gridFunc = h.newNoneGrid
	default:
		return fmt.Errorf("model: encountered unknown layout '%s' when attempting to generate hero grid", layout)
	}

	// Verify that the bracket is valid
	if _, ok := cli.Brackets[bracket]; !ok {
		return fmt.Errorf("model: encountered unknown bracket '%s' when attempting to generate hero grid", bracket)
	}

	// Make new grid
	cfgName := getCfgName(gridName, bracket)
	grid, err := gridFunc(cfgName, bracket, heroes)
	if err != nil {
		return nil
	}

	// Add grid to HeroGridConfig
	err = h.addHeroGrid(cfgName, grid)
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
func (h *HeroGridConfig) getHeroGrid(name string) *HeroGrid {
	if idx, ok := h.findGridIdx(name); ok {
		// Modify existing
		fmt.Printf("%d\n", idx)
		herogrid := h.HeroGrids[idx]
		return &herogrid
	}
	return nil
}

func (h *HeroGridConfig) newSingleGrid(gridName string, bracket string, heroes *Heroes) (*HeroGrid, error) {
	grid, err := NewHeroGrid(gridName, []string{"Heroes"})
	if err != nil {
		return nil, err
	}

	for _, hero := range *heroes {
		grid.Categories[0].HeroIDs = append(grid.Categories[0].HeroIDs, hero.HeroID)
	}

	return grid, nil
}

func (h *HeroGridConfig) newMainStatGrid(gridName string, bracket string, heroes *Heroes) (*HeroGrid, error) {
	grid, err := NewHeroGrid(gridName, []string{"Strength", "Agility", "Intellect"})
	if err != nil {
		return nil, err
	}

	var categoryIdx = map[string]int{"str": 0, "agi": 1, "int": 2}
	for _, hero := range *heroes {
		idx := categoryIdx[hero.PrimaryAttr]
		grid.Categories[idx].HeroIDs = append(grid.Categories[idx].HeroIDs, hero.HeroID)
	}

	return grid, nil
}

func (h *HeroGridConfig) newAttackTypeGrid(gridName string, bracket string, heroes *Heroes) (*HeroGrid, error) {
	grid, err := NewHeroGrid(gridName, []string{"Melee", "Ranged"})
	if err != nil {
		return nil, err
	}

	var categoryIdx = map[string]int{"melee": 0, "ranged": 1}
	for _, hero := range *heroes {
		idx := categoryIdx[hero.AttackType]
		grid.Categories[idx].HeroIDs = append(grid.Categories[idx].HeroIDs, hero.HeroID)
	}

	return grid, nil
}

func (h *HeroGridConfig) newRoleGrid(gridName string, bracket string, heroes *Heroes) (*HeroGrid, error) {
	return nil, errors.New("Role grid is not yet implemented")
}

func (h *HeroGridConfig) newLegsGrid(gridName string, bracket string, heroes *Heroes) (*HeroGrid, error) {
	return nil, errors.New("Legs grid is not yet implemented")
}

func (h *HeroGridConfig) newNoneGrid(gridName string, bracket string, heroes *Heroes) (*HeroGrid, error) {
	return nil, errors.New("None grid is not yet implemented")
}

func (h *HeroGridConfig) addHeroGrid(gridName string, grid *HeroGrid) error {
	if gridIdx, ok := h.findGridIdx(gridName); ok {
		h.HeroGrids[gridIdx] = *grid
	} else {
		h.HeroGrids = append(h.HeroGrids, *grid)
	}
	return nil
}

type HeroGrid struct {
	config.Config
	ConfigName string     `json:"config_name"`
	Categories []Category `json:"categories"`
}

func NewHeroGrid(name string, categories []string) (*HeroGrid, error) {
	var h = &HeroGrid{ConfigName: name}

	// TODO: This should probably be dynamic somehow
	switch len(categories) {
	case 1:
		h.Categories = singleCategory
	case 2:
		h.Categories = doubleCategory
	case 3:
		h.Categories = tripleCategory
	default:
		return nil, fmt.Errorf("model.NewHeroGrid: expected <=3 categories, received %d", len(categories))
	}

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
func NewCategory(name string, xpos, ypos, width, height float64) *Category {
	return &Category{
		CategoryName: name,
		XPosition:    math.Abs(xpos),
		YPosition:    math.Abs(ypos),
		Width:        math.Abs(width),
		Height:       math.Abs(height),
	}
}

// NewCategoryDefault creates a new Category with default parameters
func NewCategoryDefault(name string, xpos, ypos, width, height float64) *Category {
	return &Category{
		CategoryName: "name",
		XPosition:    0.0,
		YPosition:    0.0,
		Width:        1180.0,
		Height:       180.0,
	}
}

func (c *Category) String() string {
	var s string
	s += fmt.Sprintf("CategoryName: %s\n", c.CategoryName)
	s += fmt.Sprintf("XPosition: %f\n", c.XPosition)
	s += fmt.Sprintf("YPosition: %f\n", c.YPosition)
	s += fmt.Sprintf("Width: %f\n", c.Width)
	s += fmt.Sprintf("Height: %f\n", c.Height)
	// How can we do this with fewer allocations? Super inefficent, no?
	if len(c.HeroIDs) > 0 {
		s += "["
		for i, id := range c.HeroIDs {
			s += fmt.Sprintf("%d", id)
			if i != (len(c.HeroIDs) - 1) {
				s += ", "
			}
		}
		s += "]"
	}
	return s
}

func getCfgName(basename string, bracket string) string {
	// FIXME: This should include layout as well
	return fmt.Sprintf("%s (%s)", basename, bracket)
}

var singleCategory = []Category{
	*(NewCategory("Heroes", 0, 0, 1180, 1180)),
}

var doubleCategory = []Category{
	*(NewCategory("One", 0, 0, 1180, 280)),
	*(NewCategory("Two", 0, 300, 1180, 280)),
}

var tripleCategory = []Category{
	*(NewCategory("One", 0, 0, 1180, 180)),
	*(NewCategory("Two", 0, 200, 1180, 180)),
	*(NewCategory("Three", 0, 400, 1180, 180)),
}
