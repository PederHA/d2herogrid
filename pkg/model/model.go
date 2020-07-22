package model

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/PederHA/d2herogrid/pkg/config"
)

type HeroGridConfig struct {
	config.Config
	Version int      `json:"version"`
	Configs []Config `json:"configs"`
}

type Config struct {
	config.Config
	ConfigName string     `json:"config_name"`
	Categories []Category `json:"categories"`
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
