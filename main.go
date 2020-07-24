package main

import (
	"log"

	"github.com/PederHA/d2herogrid/cmd"
	"github.com/PederHA/d2herogrid/cmd/cli"
	"github.com/PederHA/d2herogrid/pkg/model"
)

func main() {
	cfg, err := cli.Parse()
	if err != nil {
		panic(err)
	}
	hgc, err := model.NewHeroGridConfig(cfg.Path)
	if err != nil {
		panic(err)
	}

	app := cmd.NewApp(cfg, hgc)

	err = app.Run()
	if err != nil {
		log.Fatal(err)
	}
}
