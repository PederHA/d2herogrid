package model

import (
	"io"
	"os"
	"sort"
	"testing"

	_ "github.com/PederHA/d2herogrid/testing"
)

const (
	heroesPath = "testing/heroes.json"
	hgcPath    = "testing/hero_grid_config.json"
	hgcPathBak = hgcPath + ".bak"
)

var gridNames = []string{"Custom Test", "Mainstat Test"}

var heroes *Heroes
var hgc *HeroGridConfig

func init() {
	var err error

	heroes, err = NewHeroesFromFile(heroesPath)
	if err != nil {
		panic(err)
	}

	hgc, err = NewHeroGridConfig(hgcPath)
	if err != nil {
		panic(err)
	}

	err = copyFile(hgcPath, hgcPathBak)
	if err != nil {
		panic(err)
	}

}

func teardown() {
	err := copyFile(hgcPathBak, hgcPath)
	if err != nil {
		panic(err)
	}
}

func copyFile(srcPath, destPath string) error {
	inputFile, err := os.Open(srcPath)
	if err != nil {
		return err
	}

	outputFile, err := os.Create(destPath)
	if err != nil {
		return err
	}

	defer outputFile.Close()
	_, err = io.Copy(outputFile, inputFile)
	inputFile.Close()
	if err != nil {
		return err
	}
	return nil
}

func TestMain(m *testing.M) {
	// Set-up ...
	os.Exit(m.Run())
	teardown()
}

func TestHeroGridConfig_MakeGrid_LayoutModify(t *testing.T) {
	for _, gridName := range gridNames {
		for _, bracket := range AllBrackets {
			heroes.SetSorting(bracket)
			sort.Sort(heroes)
			err := hgc.MakeGrid(gridName, LayoutModify, bracket, heroes)
			if err != nil {
				t.Error(err)
			}
		}
	}
}

// TestHeroGridConfig_MakeGrid tests HeroGridConfig.MakeGrid(...) with all brackets
// and all layouts except LayoutModify
func TestHeroGridConfig_MakeGrid(t *testing.T) {
	for _, bracket := range AllBrackets {
		heroes.SetSorting(bracket)
		sort.Sort(heroes)
		for _, layout := range AllLayouts {
			for _, gridName := range gridNames {
				err := hgc.MakeGrid(gridName, layout, bracket, heroes)
				if err != nil {
					t.Error(err)
				}
			}
		}
	}
}
