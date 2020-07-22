package client

import (
	"encoding/json"
	"net/http"
)

type heroes []int

func (h *heroes) Sort(ascending bool) {

}

func Get() (*Heroes, error) {
	r, err := http.Get("https://api.opendota.com/api/heroStats")
	if err != nil {
		return nil, err
	}
	var h = &Heroes{}
	var heroes = []HeroStats{}
	err = json.NewDecoder(r.Body).Decode(&heroes)
	if err != nil {
		return nil, err
	}
	h.Heroes = heroes
	return h, nil

}
