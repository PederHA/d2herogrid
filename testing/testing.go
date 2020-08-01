// Copied from https://brandur.org/fragments/testing-go-project-root
// When this module is imported, the current working directory is set to
// the project's root directory, as opposed to the directory of the package

package testing

import (
	"os"
	"path"
	"runtime"
)

func init() {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "..")
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}
}
