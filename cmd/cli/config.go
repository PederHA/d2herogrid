// The cli package implements

package cli

import (
	"fmt"
	"log"
	"runtime"

	"github.com/PederHA/d2herogrid/pkg/config"
)

type UserConfig struct {
	config.Config
	Brackets      []int  `json:"brackets" yaml:"brackets"`
	GridName      string `json:"grid_name" yaml:"grid_name"`
	Layout        int    `json:"layout" yaml:"layout"`
	Path          string `json:"path" yaml:"path"`
	SortAscending bool   `json:"sort_ascending" yaml:"sort_ascending"`
}

func NewUserConfig(brackets []int, gridName string, layout int, path string, sortAsc bool) *UserConfig {
	return &UserConfig{
		Brackets:      brackets,
		GridName:      gridName,
		Layout:        1,
		Path:          path,
		SortAscending: sortAsc,
	}
}

func NewUserConfigDefaults() *UserConfig {
	path, err := autodetectUserdataDir()
	if err != nil {
		log.Println(err)
	}
	// TODO: Do something like Python's Path.iterdir() to get subdirectories
	// 		 so user can choose the correct userdata directory
	return NewUserConfig(
		defaultBrackets,
		defaultGridName,
		defaultLayout,
		path,
		defaultSortAscending,
	)
}

func autodetectUserdataDir() (string, error) {
	switch runtime.GOOS {
	case "windows":
		return "C:/Program Files(x86)/Steam/userdata/", nil
	case "darwin":
		return "/Users/Peder-MAC/Library/Application Support/Steam/userdata/19123403/570/remote/cfg/hero_grid_config.json", nil
		//return "~/Library/Application Support/Steam/userdata/", nil
	case "linux":
		return "~/Steam/userdata/", nil
	default:
		return "", fmt.Errorf("config: Unknown OS")
	}
}
