package model

type HeroGridConfig struct {
	Version int      `json:"version"`
	Configs []Config `json:"configs"`
}

type Config struct {
	ConfigName string     `json:"config_name"`
	Categories []Category `json:"categories"`
}

type Category struct {
	CategoryName string  `json:"category_name"`
	XPosition    float32 `json:"x_position"`
	YPosition    float32 `json:"y_position"`
	Width        float32 `json:"width"`
	Height       float32 `json:"height"`
	HeroIDs      []int   `json:"hero_ids"` // uint8?
}
