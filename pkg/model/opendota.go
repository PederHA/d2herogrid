package model

import "fmt"

//type Heroes struct {
//	Heroes []HeroStats
//}

type Heroes []*HeroStats

func (h *Heroes) SetSorting(bracket string) {
	for _, hero := range *h {
		hero.setSorting(bracket)
	}
}

func (h *Heroes) PrintWinrates() {
	for _, hero := range *h {
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
	SortingWin    int    //`json:"8_win"`
	SortingPick   int    //`json:"8_pick"`
}

func (h *HeroStats) setSorting(bracket string) {
	// Use a map or something instead?
	// Or some stupid meta-programming with reflection?
	switch bracket {
	case "Herald":
		h.SortingWin = h.HeraldWin
		h.SortingPick = h.HeraldPick
	case "Guardian":
		h.SortingWin = h.GuardianWin
		h.SortingPick = h.GuardianPick
	case "Crusader":
		h.SortingWin = h.CrusaderWin
		h.SortingPick = h.CrusaderPick
	case "Archon":
		h.SortingWin = h.ArchonWin
		h.SortingPick = h.ArchonPick
	case "Legend":
		h.SortingWin = h.LegendWin
		h.SortingPick = h.LegendPick
	case "Ancient":
		h.SortingWin = h.AncientWin
		h.SortingPick = h.AncientPick
	case "Divine":
		h.SortingWin = h.DivineWin
		h.SortingPick = h.DivinePick
	case "Immortal":
		h.SortingWin = h.ImmortalWin
		h.SortingPick = h.ImmortalPick
	case "Pro":
		h.SortingWin = h.ProWin
		h.SortingPick = h.ProPick
	}
	// Avoid division by 0 error
	if h.SortingPick < 1 {
		h.SortingPick = 1
	}
}
