package cmd

import (
	"sort"

	"github.com/PederHA/d2herogrid/cmd/cli"
	"github.com/PederHA/d2herogrid/cmd/client"
	"github.com/PederHA/d2herogrid/pkg/model"
)

type App struct {
	UserConfig     *cli.UserConfig
	HeroGridConfig *model.HeroGridConfig
}

// NewApp creates a new d2herogrid app
func NewApp(config *cli.UserConfig, hgc *model.HeroGridConfig) *App {
	return &App{
		UserConfig:     config,
		HeroGridConfig: hgc,
	}
}

// Run fetches OpenDota hero data, then creates HeroGrids as specified in the UserConfig
func (a *App) Run() error {
	// Get Hero Winrates from OpenDota API
	heroes, err := client.GetHeroStats()
	if err != nil {
		return err
	}

	// Make hero grids
	for _, bracket := range a.UserConfig.Brackets {
		a.makeGrid(bracket, heroes)
	}

	// Save hero grids
	err = a.HeroGridConfig.SaveConfigJSON(a.UserConfig.Path)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) makeGrid(bracket string, heroes *model.Heroes) error {
	heroes.SetSorting(bracket) // FIX
	heroes.PrintWinrates()     // REMOVE

	// Sort heroes by winrate
	if a.UserConfig.SortAscending {
		sort.Sort(heroes)
	} else {
		sort.Sort(sort.Reverse(heroes))
	}

	heroes.PrintWinrates() // REMOVE

	// Create New Hero Grid using specified layout
	err := a.HeroGridConfig.MakeGrid(a.UserConfig.GridName, a.UserConfig.Layout, bracket, heroes)
	if err != nil {
		return err
	}

	return nil
}
