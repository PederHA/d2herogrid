package cmd

import (
	"fmt"
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
	// Sort heroes by winrate
	// TODO: Winrate in a specific skill bracket
	sort.Sort(model.ByDivine(heroes.Heroes))
	fmt.Printf("%v", heroes)

	// Create New Hero Grid using specified layout
	//

	// Save hero grid
	//

	// No error occured
	return nil
}
