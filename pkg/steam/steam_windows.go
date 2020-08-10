// +build windows

package steam

import "golang.org/x/sys/windows/registry"

func Path() (string, error) {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\WOW6432Node\Valve\Steam`, registry.QUERY_VALUE)
	defer k.Close()
	if err != nil {
		return "", err
	}

	s, _, err := k.GetStringValue("InstallPath")
	if err != nil {
		return "", nil // NOTE: Should we not propagate this error?
	}

	return s, nil
}
