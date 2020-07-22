package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/PederHA/d2herogrid/pkg/config"
)

type HeroGridConfig struct {
	config.Config
	Version int      `json:"version"`
	Configs []Config `json:"configs"`
}

func (h *HeroGridConfig) String() string {
	var s string
	s += fmt.Sprintf("Version: %d\n", h.Version)
	for _, config := range h.Configs {
		s += config.String()
	}
	return s
}

type Config struct {
	config.Config
	ConfigName string     `json:"config_name"`
	Categories []Category `json:"categories"`
}

func (c *Config) String() string {
	var s string
	s += fmt.Sprintf("ConfigName: %s\n", c.ConfigName)
	for _, category := range c.Categories {
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

func NewHeroGridConfig(hgcPath string) (*HeroGridConfig, error) {
	var hgc = &HeroGridConfig{}
	f, err := os.Open("hero_grid_config.json")
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

func (h *HeroGridConfig) ListGrids() {
	fmt.Printf("Grids:\n")
	for _, config := range h.Configs {
		fmt.Printf("\tConfig Name: %s\n", config.ConfigName)
		for _, category := range config.Categories {
			fmt.Printf("\t\tCategory Name: %s\n", category.CategoryName)
		}
	}
}
