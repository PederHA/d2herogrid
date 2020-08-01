// TODO: Heroes should probably be called HeroWinrates or something along those lines
// Or called HeroStats, and contain []Hero

package model

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/PederHA/d2herogrid/internal/utils"
)

// Heroes is a slice of HeroStats pointers
type Heroes []*HeroStats

// SetSorting specifies a specific skill bracket that heroes should be sorted by.
func (h *Heroes) SetSorting(bracket *Bracket) {
	for _, hero := range *h {
		hero.setSorting(bracket)
	}
}

func (h *Heroes) DumpJSON(path string) error {
	return utils.MarshalJSON(path, h)
}

// NewHeroesFromAPI constructs a new Heroes object from an OpenDota API call
func NewHeroesFromAPI() (*Heroes, error) {
	r, err := http.Get("https://api.opendota.com/api/heroStats")
	if err != nil {
		return nil, err
	}
	return newHeroes(r.Body)
}

// NewHeroesFromFile creates a new Heroes object by unmarshaling a JSON-encoded file
func NewHeroesFromFile(filePath string) (*Heroes, error) {
	r, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	return newHeroes(r)
}

// newHeroes is an internal JSON decoder function that decodes JSON-encoded data
// from an io.Reader, used for both NewHeroesFrom... functions
func newHeroes(r io.Reader) (*Heroes, error) {
	heroes := new(Heroes)
	err := json.NewDecoder(r).Decode(heroes)
	if err != nil {
		return nil, err
	}
	return heroes, nil
}

// Unused + internal.
// FIXME: delete?!
func (h *Heroes) printWinrates() {
	for _, hero := range *h {
		if hero.SortingWin == 0 {
			hero.setSorting(BracketDivine)
		}
		fmt.Printf("%f ", float64(hero.SortingWin)/float64(hero.SortingPick))
	}
	fmt.Print("\n")
}

func (h Heroes) Len() int {
	return len(h)
}

func (h Heroes) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h Heroes) Less(i, j int) bool {
	return (float64(h[i].SortingWin) / float64(h[i].SortingPick)) <
		(float64(h[j].SortingWin) / float64(h[j].SortingPick))
}

// HeroStats represents the stats of a single Dota 2 hero
type HeroStats struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	LocalizedName string `json:"localized_name"`
	Img           string `json:"img"`
	Icon          string `json:"Icon"`
	HeroID        int    `json:"hero_id"`
	PrimaryAttr   string `json:"primary_attr"`
	AttackType    string `json:"attack_type"`
	Legs          int    `json:"legs"`
	HeraldWin     int    `json:"1_win"`
	HeraldPick    int    `json:"1_pick"`
	GuardianWin   int    `json:"2_pick"`
	GuardianPick  int    `json:"2_win"`
	CrusaderWin   int    `json:"3_win"`
	CrusaderPick  int    `json:"3_pick"`
	ArchonWin     int    `json:"4_win"`
	ArchonPick    int    `json:"4_pick"`
	LegendWin     int    `json:"5_win"`
	LegendPick    int    `json:"5_pick"`
	AncientWin    int    `json:"6_win"`
	AncientPick   int    `json:"6_pick"`
	DivineWin     int    `json:"7_win"`
	DivinePick    int    `json:"7_pick"`
	ImmortalWin   int    `json:"8_win"`
	ImmortalPick  int    `json:"8_pick"`
	ProWin        int    `json:"pro_win"`
	ProPick       int    `json:"pro_pick"`
	ProBan        int    `json:"pro_ban"`
	SortingWin    int    `json:"-"`
	SortingPick   int    `json:"-"`
}

func (h *HeroStats) setSorting(bracket *Bracket) {
	// Use a map or something instead?
	// Or some stupid meta-programming with reflection?
	switch bracket {
	case BracketHerald:
		h.SortingWin = h.HeraldWin
		h.SortingPick = h.HeraldPick
	case BracketGuardian:
		h.SortingWin = h.GuardianWin
		h.SortingPick = h.GuardianPick
	case BracketCrusader:
		h.SortingWin = h.CrusaderWin
		h.SortingPick = h.CrusaderPick
	case BracketArchon:
		h.SortingWin = h.ArchonWin
		h.SortingPick = h.ArchonPick
	case BracketLegend:
		h.SortingWin = h.LegendWin
		h.SortingPick = h.LegendPick
	case BracketAncient:
		h.SortingWin = h.AncientWin
		h.SortingPick = h.AncientPick
	case BracketDivine:
		h.SortingWin = h.DivineWin
		h.SortingPick = h.DivinePick
	case BracketImmortal:
		h.SortingWin = h.ImmortalWin
		h.SortingPick = h.ImmortalPick
	case BracketPro:
		h.SortingWin = h.ProWin
		h.SortingPick = h.ProPick
	}
	// Avoid division by 0 error
	if h.SortingPick < 1 {
		h.SortingPick = 1
	}
}
