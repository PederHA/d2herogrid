package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/PederHA/d2herogrid/cmd/cli"
	"github.com/PederHA/d2herogrid/pkg/model"
)

func read() (*model.HeroGridConfig, error) {
	var hgc = &model.HeroGridConfig{}

	f, err := os.Open("hero_grid_config.json")
	defer f.Close()
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(b, hgc)

	return hgc, nil
}

func main() {

	_, err := read()
	if err != nil {
		log.Fatal(err)
	}

	c := &cli.Config{
		Brackets:      []int{7, 8},
		DefaultName:   "OpenDotaHeroWinrates",
		Layout:        1,
		Path:          "some/path/lol",
		SortAscending: false}

	err = c.Set("Brackets", []int{1, 2, 3, 4})
	if err != nil {
		log.Fatal(err)
	}

	b, err := c.Get("Brackets")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Brackets: %v\n", b)
	//fmt.Printf("%v\n", hgc)
}
