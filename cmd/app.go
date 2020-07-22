package cmd

import (
	"fmt"
	"log"
	"sort"

	"github.com/PederHA/d2herogrid/cmd/cli"
	"github.com/PederHA/d2herogrid/cmd/client"
	"github.com/PederHA/d2herogrid/pkg/model"
)

type App struct {
	UserConfig     *cli.UserConfig
	HeroGridConfig *model.HeroGridConfig
}

func NewApp(config *cli.UserConfig, hgc *model.HeroGridConfig) *App {
	return &App{
		UserConfig:     config,
		HeroGridConfig: hgc,
	}
}

func (a *App) Run() {
	// do stuff
	//a.HeroGridConfig.ListGrids()
	heroes, err := client.Get()
	if err != nil {
		log.Fatal(err)
	}
	sort.Sort(client.ByDivine(heroes.Heroes))
	fmt.Printf("%v", heroes)
}
