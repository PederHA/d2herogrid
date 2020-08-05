package cli

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/PederHA/d2herogrid/internal/utils"
	"github.com/PederHA/d2herogrid/pkg/model"
)

var (
	errNoValidBrackets = errors.New("parse: no valid brackets given")
	invalidLayout      = "invalid layout '%s'"
)

// Parse parses command-line arguments and returns a new UserConfig
func Parse(name, layout, path *string, sortAsc *bool, brackets []string) (*UserConfig, error) {
	// LAYOUT
	l, err := parseLayout(layout)
	if err != nil {
		return nil, err
	}

	// PATH
	p, err := parsePath(path)
	if err != nil {
		return nil, err
	}

	// BRACKETS
	// Add default brackets if no args are given
	if len(brackets) == 0 { // NOTE: is this the correct comparison?
		for _, b := range DefaultBrackets {
			brackets = append(brackets, b.Aliases[0])
		}
	}
	b, err := parseBrackets(brackets)
	if err != nil {
		return nil, err
	}

	return NewUserConfig(
		*name,
		b,
		l,
		*p,
		*sortAsc,
	), nil
}

func parseBrackets(br []string) (model.Brackets, error) {
	var validBrackets model.Brackets
	keys := make(map[*model.Bracket]bool) // Avoid duplicates

	for _, b := range br { // For each argument
		b = strings.ToLower(b)
	bracketLoop:
		for _, bracket := range model.AllBrackets {
			// Treat formatted name as an alias as well
			aliases := append(bracket.Aliases, bracket.Name)
			for _, alias := range aliases {
				if _, ok := keys[bracket]; !ok && b == alias { // if arg == an alias
					keys[bracket] = true
					validBrackets = append(validBrackets, bracket)
					break bracketLoop
				}
			}
		}
	}

	if len(validBrackets) == 0 {
		return nil, errNoValidBrackets
	}

	return validBrackets, nil
}

func parseLayout(layout *string) (*model.Layout, error) {
	for _, l := range model.AllLayouts {
		// Treat formatted name as an alias as well
		aliases := append(l.Aliases, l.Name)
		for _, alias := range aliases {
			if *layout == alias {
				return l, nil
			}
		}
	}
	return nil, fmt.Errorf(invalidLayout, *layout)
}

func parsePath(path *string) (*string, error) {
	// TODO: Ask user for subdirectory in userdata
	if *path == "" {
		p, err := autodetectUserdataDir()
		if err != nil {
			return nil, err
		}
		*path = p
	}

	// Check if directory exists
	err := utils.CheckDirExists(*path)
	if err != nil {
		return nil, err
	}

	// Create filepath for hero_grid_config.json
	fp := filepath.Join(*path, "hero_grid_config.json")
	err = utils.CheckFileExists(fp)
	if err != nil {
		// Make grid if it doesn't exist
		if os.IsNotExist(err) {
			fmt.Printf("Creating new hero grid config at '%s'\n", fp)
			hgc := model.NewHeroGridConfigDefault()
			err := hgc.SaveConfigJSON(fp)
			if err != nil {
				fmt.Printf("Failed to create hero grid config at '%s'\n", fp)
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	return &fp, nil
}
