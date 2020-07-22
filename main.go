package main

import (
	"github.com/PederHA/d2herogrid/cmd"
	"github.com/PederHA/d2herogrid/cmd/cli"
	"github.com/PederHA/d2herogrid/pkg/model"
)

func main() {
	cfg := cli.Parse()
	hgc, err := model.NewHeroGridConfig(cfg.Path)
	if err != nil {
		panic(err)
	}

	app := cmd.NewApp(cfg, hgc)

	app.Run()
}
