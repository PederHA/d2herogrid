// Platform-independent test

package steam

import (
	"testing"
)

func TestPath(t *testing.T) {
	p, err := Path()
	if err != nil {
		t.Error(err)
	}
	if p == "" {
		t.Errorf("Path() returned empty string and no error.")
	}
}
