package client

import "fmt"

type Heroes struct {
	Heroes []HeroStats
}

type ByDivine []HeroStats

func (h ByDivine) Len() int      { return len(h) }
func (h ByDivine) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h ByDivine) Less(i, j int) bool {
	f := func(pick int) int {
		if pick < 0 {
			return 1
		}
		return pick
	}
	iPick := f(h[i].DivinePick)
	jPick := f(h[j].DivinePick)
	res := (h[i].DivineWin / iPick) < (h[j].DivineWin / jPick)
	if res {
		fmt.Printf("LOL!\n")
	}
	return res
}

type HeroStats struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	LocalizedName string `json:"localized_name"`
	Img           string `json:"img"`
	Icon          string `json:"Icon"`
	HeroID        int    `json:"hero_id"`
	ProWin        int    `json:"pro_win"`
	ProPick       int    `json:"pro_pick"`
	ProBan        int    `json:"pro_ban"`
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
}
