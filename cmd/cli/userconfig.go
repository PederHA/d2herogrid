// The cli package implements

package cli

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/PederHA/d2herogrid/internal/utils"
	"github.com/PederHA/d2herogrid/pkg/config"
	"gopkg.in/yaml.v3"
)

type UserConfig struct {
	config.Config `json:"-" yaml:"-"`
	Brackets      []string `json:"brackets" yaml:"brackets"`
	GridName      string   `json:"grid_name" yaml:"grid_name"`
	Layout        string   `json:"layout" yaml:"layout"`
	Path          string   `json:"path" yaml:"path"`
	SortAscending bool     `json:"sort_ascending" yaml:"sort_ascending"`
}

func NewUserConfig(brackets []string, gridName string, layout string, path string, sortAsc bool) *UserConfig {
	return &UserConfig{
		Brackets:      brackets,
		GridName:      gridName,
		Layout:        layout,
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

func (u *UserConfig) DumpYaml() error {
	cfgFname := "config.yaml"

	err := u.createConfigDir(userconfigDir)
	if err != nil {
		return err
	}

	b, err := yaml.Marshal(u)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filepath.Join(userconfigDir, cfgFname), b, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserConfig) createConfigDir(dirname string) error {
	if dirname == "" {
		return fmt.Errorf("config: filepath cannot be empty")
	}

	// TODO: Handle malformed paths + directory paths

	// Check if config dir exists
	if ok, err := utils.CheckDirExists(dirname); err != nil {
		return err
	} else if !ok {
		os.MkdirAll(dirname, os.ModePerm)
	}
	return nil

	//// Check if config file exists
	//if ok, err := utils.CheckFileExists(fullPath); err != nil {
	//	return err
	//} else if !ok {
	//	f, err := os.Create(fullPath)
	//	defer f.Close()
	//	if err != nil {
	//		return err
	//	}
	//}

	return nil
}

func autodetectUserdataDir() (string, error) {
	switch runtime.GOOS {
	case "windows":
		return "D:\\Programming\\Go\\src\\d2herogrid\\hero_grid_config.json", nil
		//return "C:\\Program Files (x86)\\Steam\\userdata\\19123403\\570\remote\\cfg\\hero_grid_config.json", nil
		//return "C:/Program Files(x86)/Steam/userdata/", nil
	case "darwin":
		return "/Users/Peder-MAC/Library/Application Support/Steam/userdata/19123403/570/remote/cfg/hero_grid_config.json", nil
		//return "~/Library/Application Support/Steam/userdata/", nil
	case "linux":
		return "~/Steam/userdata/", nil
	default:
		return "", fmt.Errorf("config: Unknown OS")
	}
}
