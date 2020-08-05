package cli

import (
	"fmt"
	"testing"

	"github.com/PederHA/d2herogrid/pkg/model"
	_ "github.com/PederHA/d2herogrid/testing"
)

func TestAutodetectUserdataDir(t *testing.T) {
	d, err := autodetectUserdataDir()
	if err != nil {
		t.Error(err)
	}
	if d == "userdata" {
		t.Error("autodetectUserdataDir() failed, got 'userdata', want '<absolute_path_to_steam_directory>/userdata'")
	}
}

func TestNewUserConfigDefaults(t *testing.T) {
	got := NewUserConfigDefaults()
	want := UserConfig{
		GridName:      DefaultGridName,
		Brackets:      DefaultBrackets,
		Layout:        DefaultLayout,
		SortAscending: DefaultSortAscending,
	}

	// NOTE: NOT testing value of UserConfig.Path

	if got.GridName != want.GridName {
		t.Errorf("NewUserConfigDefaults() failed, want name '%s', got '%s'", want.GridName, got.GridName)
	}
	for _, gb := range got.Brackets {
		for _, wb := range want.Brackets {
			if gb.Name != wb.Name {
				t.Errorf("NewUserConfigDefaults() failed, want bracket '%s', got bracket '%s'", wb.Name, gb.Name)
			}
		}
	}
	if got.Layout.Name != want.Layout.Name {
		t.Errorf("NewUserConfigDefaults() failed, want layout '%s', got layout '%s'", want.Layout.Name, got.Layout.Name)
	}

	if got.SortAscending != want.SortAscending {
		t.Errorf("NewUserConfigDefaults() failed, want SortAscending '%t', got '%t'", want.SortAscending, got.SortAscending)
	}
}

func TestNewUserConfigFromFile(t *testing.T) {
	want := &UserConfig{
		GridName:      "d2hg",
		Brackets:      model.Brackets{model.BracketDivine, model.BracketImmortal},
		Layout:        model.LayoutMainStat,
		Path:          "hero_grid_config.json",
		SortAscending: false,
	}

	cfgPath := "testing/config.yaml"
	got, err := NewUserConfigFromFile(cfgPath)
	if err != nil {
		t.Error(err)
	}

	err = compareUserConfigs(got, want)
	if err != nil {
		t.Error(err)
	}
}

func TestUserConfigDumpYAML(t *testing.T) {
	testconf := "testing/config_test.yaml"
	uc := &UserConfig{
		GridName:      "d2hg",
		Brackets:      model.Brackets{model.BracketDivine, model.BracketImmortal},
		Layout:        model.LayoutMainStat,
		Path:          "hero_grid_config.json",
		SortAscending: false,
	}

	err := uc.DumpYAML(testconf)
	if err != nil {
		t.Error(err)
	}

	got, err := NewUserConfigFromFile(testconf)
	if err != nil {
		t.Error(err)
	}

	err = compareUserConfigs(got, uc)
	if err != nil {
		t.Error(err)
	}

}

func compareUserConfigs(got, want *UserConfig) error {
	// NAME
	if got.GridName != want.GridName {
		return fmt.Errorf(
			"TestNewUserConfigFromFile() failed. got GridName == %q, want %q",
			got.GridName, want.GridName)
	}

	// BRACKETS
	if len(got.Brackets) != len(want.Brackets) {
		return fmt.Errorf(
			"TestNewUserConfigFromFile() failed. Mismatched brackets length. Want %d, got %d",
			len(want.Brackets), len(got.Brackets))
	}

	for i := range got.Brackets {
		if got.Brackets[i] != want.Brackets[i] {
			return fmt.Errorf(
				"TestNewUserConfigFromFile() failed. Got Brackets[%d] == %q, want Brackets[%d] == %q",
				i, got.Brackets[i].Name, i, want.Brackets[i].Name)
		}
	}

	// LAYOUT
	if got.Layout.Name != want.Layout.Name {
		return fmt.Errorf("TestNewUserConfigFromFile() failed. Got LayoutName == %q, want %q",
			got.Layout.Name, want.Layout.Name)
	}

	// PATH
	if got.Path != want.Path {
		return fmt.Errorf("TestNewUserConfigFromFile() failed. Got Path == %q, want %q",
			got.Path, want.Path)
	}

	// SORT
	if got.SortAscending != want.SortAscending {
		return fmt.Errorf("TestNewUserConfigFromFile() failed. Got SortAscending == %t, want %t",
			got.SortAscending, want.SortAscending)
	}

	return nil
}
