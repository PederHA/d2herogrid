// +build windows

package cli

import "golang.org/x/sys/windows/registry"

func getSteamPathWindows() (string, error) {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\WOW6432Node\Valve\Steam`, registry.QUERY_VALUE)
	if err != nil {
		return "", err
	}
	defer k.Close()

	s, _, err := k.GetStringValue("InstallPath")
	if err != nil {
		return "", nil // NOTE: Should we not propagate this error?
	}
	return s, nil
}
