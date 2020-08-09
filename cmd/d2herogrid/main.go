// go: generate goversioninfo -icon=assets/logo.ico

package main

import (
	"flag"
	"log"

	"github.com/PederHA/d2herogrid/pkg/cli"
	"github.com/PederHA/d2herogrid/pkg/model"
)

var (
	name     *string
	layout   *string
	path     *string
	sortAsc  *bool
	brackets []string
)

func init() {
	// CLI parameters
	name = flag.String("n", cli.DefaultGridName, "Grid name")
	layout = flag.String("l", cli.DefaultLayout.Aliases[0], "Grid layout")
	path = flag.String("p", ".", "Path to Dota 2 userdata directory")
	sortAsc = flag.Bool("s", false, "Sort ascending (low-high) [default: high-low]")
	flag.Parse()
	brackets = flag.Args()
}

func run() error {
	// Parse CLI args
	cfg, err := cli.Parse(name, layout, path, sortAsc, brackets)
	if err != nil {
		return err
	}

	// Load contents of hero_grid_config.json
	hgc, err := model.NewHeroGridConfig(cfg.Path)
	if err != nil {
		return err
	}

	// Save user config (Implemented, but useless)
	//err = cfg.DumpYAML(cli.ConfigPath)
	//if err != nil {
	//	return err
	//}

	app := NewApp(cfg, hgc)

	err = app.Run()
	if err != nil {
		return err
	}

	return nil
}

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}
