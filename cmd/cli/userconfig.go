// The cli package implements

package cli

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"

	"github.com/PederHA/d2herogrid/internal/utils"
	"github.com/PederHA/d2herogrid/pkg/config"
	"github.com/PederHA/d2herogrid/pkg/model"
	"gopkg.in/yaml.v3"
)

type UserConfig struct {
	config.Config `json:"-" yaml:"-"`
	GridName      string         `json:"grid_name" yaml:"grid_name"`
	Brackets      model.Brackets `json:"brackets" yaml:"brackets"`
	Layout        *model.Layout  `json:"layout" yaml:"layout"`
	Path          string         `json:"path" yaml:"path"`
	SortAscending bool           `json:"sort_ascending" yaml:"sort_ascending"`
}

func NewUserConfig(gridName string, brackets model.Brackets, layout *model.Layout, path string, sortAsc bool) *UserConfig {
	return &UserConfig{
		GridName:      gridName,
		Brackets:      brackets,
		Layout:        layout,
		Path:          path,
		SortAscending: sortAsc,
	}
}

// NewUserConfigDefaults creates a new UserConfig using default parameters,
// which are documented in defaults.go
func NewUserConfigDefaults() *UserConfig {
	path, err := autodetectUserdataDir()
	if err != nil {
		log.Println(err)
	}
	// TODO: Do something like Python's Path.iterdir() to get subdirectories
	// 		 so user can choose the correct userdata directory
	return NewUserConfig(
		DefaultGridName,
		DefaultBrackets,
		DefaultLayout,
		path,
		DefaultSortAscending,
	)
}

// DumpYaml saves a UserConfig as a yaml-formatted file.
// File location is currently not very flexible, which is pretty bad.
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
	err := utils.CheckDirExists(dirname)
	if err != nil {
		if os.IsNotExist(err) {
			os.MkdirAll(dirname, os.ModePerm)
		} else {
			return err
		}
	}
	return nil
}

// FIXME: The fucking state of this thing
func autodetectUserdataDir() (string, error) {
	var s string
	var err error
	switch runtime.GOOS {
	case "windows":
		s, err = getSteamPathWindows()
		if err != nil {
			return "", err // TODO: write a more user-friendly error
		}
	case "darwin":
		// Dynamic steam path detection NYI
		s = "~/Library/Application Support/Steam"
	case "linux":
		// Dynamic steam path detection NYI
		s = "~/Steam"
	default:
		return "", fmt.Errorf("config: Unsupported OS")
	}
	return path.Join(s, "userdata"), nil
}
