// This file implements UserConfig, which holds the parsed values of the user's
// command-line arguments. Saving and loading the config as YAML is supported,
// but the program does not currently use the feature.

package cli

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"

	"github.com/PederHA/d2herogrid/internal/utils"
	"github.com/PederHA/d2herogrid/pkg/model"
	"github.com/PederHA/d2herogrid/pkg/steam"
	"github.com/mitchellh/go-homedir"
)

const (
	configFilename  = "config.yaml" // Unused
	configParentDir = ".d2herogrid"
)

var configDir string  // Absolute path of config's parent directory
var ConfigPath string // Absolute path of config

func init() {
	// FIXME: This is TERRIBLE
	dir, err := homedir.Dir()
	if err != nil {
		// NOTE: Not panicing here
		fmt.Printf("Unable to detect home directory! A persistent config cannot be saved/loaded.")
	}
	configDir = filepath.Join(dir, configParentDir)
	ConfigPath = filepath.Join(configDir, configFilename)
}

type UserConfig struct {
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

// NewUserConfigFromFile creates a new UserConfig from contents of a
// YAML-encoded config file
func NewUserConfigFromFile(path string) (*UserConfig, error) {
	uc := new(UserConfig)
	err := utils.UnmarshalYAML(path, uc)
	if err != nil {
		return nil, err
	}

	// Parse brackets value
	if len(uc.Brackets) > 0 {
		var br []string
		for _, b := range uc.Brackets {
			br = append(br, b.Name)
		}
		uc.Brackets, err = parseBrackets(br)
		if err != nil {
			return nil, err
		}
	}

	// Parse layout value
	uc.Layout, err = parseLayout(&uc.Layout.Name)
	if err != nil {
		return nil, err
	}

	return uc, nil
}

// DumpYAML saves a UserConfig as a yaml-formatted file.
// File location is currently not very flexible, which is pretty bad.
// NOTE: Change name to SaveToFile, SaveYAML, or simply Save?
func (u *UserConfig) DumpYAML(path string) error {
	err := u.createConfigDir(configDir)
	if err != nil {
		return err
	}
	return utils.MarshalYAML(path, u)
}

func (u *UserConfig) createConfigDir(dirname string) error {
	if dirname == configParentDir {
		return fmt.Errorf("config: Unable to save config when home directory cannot be detected")
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
	case "windows", "darwin", "linux":
		s, err = steam.Path()
	default:
		s, err = "", fmt.Errorf("config: Unsupported OS")
	}

	if err != nil {
		return "", err
	}

	return path.Join(s, "userdata"), nil
}
