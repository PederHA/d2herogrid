package cli

import "testing"

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
