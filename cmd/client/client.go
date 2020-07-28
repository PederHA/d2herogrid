package client

import (
	"encoding/json"
	"net/http"

	"github.com/PederHA/d2herogrid/pkg/model"
)

func GetHeroStats() (*model.Heroes, error) {
	// TODO Should it return []HeroStats?
	r, err := http.Get("https://api.opendota.com/api/heroStats")
	if err != nil {
		return nil, err
	}
	//var h = &model.Heroes{}
	//var heroes = []model.HeroStats{}
	var heroes = new(model.Heroes)
	err = json.NewDecoder(r.Body).Decode(&heroes)
	if err != nil {
		return nil, err
	}
	//h = Heroes(heroes)
	return heroes, nil

}
