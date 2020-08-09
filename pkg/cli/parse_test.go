// These tests are nowhere near as robust as the ones I wrote for the Python version.
// Mostly because I am awful at Go testing, but I do miss Pytest...

package cli

import (
	"fmt"
	"testing"

	"github.com/PederHA/d2herogrid/pkg/model"
)

type parseString struct {
	input   *string
	wantRtn *string
	wantErr error
}

type parseStringSlice struct {
	input   []string
	wantRtn []string
	wantErr error
}

type Brackets model.Brackets

var (
	herald   = model.BracketHerald
	guardian = model.BracketGuardian
	crusader = model.BracketCrusader
	archon   = model.BracketArchon
	legend   = model.BracketLegend
	ancient  = model.BracketAncient
	divine   = model.BracketDivine
	immortal = model.BracketImmortal
	pro      = model.BracketPro
)

func TestParseBrackets(t *testing.T) {
	iw := []struct {
		input []string
		want  Brackets
		err   error
	}{
		{[]string{"herald"}, Brackets{herald}, nil},
		{[]string{"guardian"}, Brackets{guardian}, nil},
		{[]string{"crusader"}, Brackets{crusader}, nil},
		{[]string{"archon"}, Brackets{archon}, nil},
		{[]string{"legend"}, Brackets{legend}, nil},
		{[]string{"ancient"}, Brackets{ancient}, nil},
		{[]string{"divine"}, Brackets{divine}, nil},
		{[]string{"immortal"}, Brackets{immortal}, nil},
		{[]string{"pro"}, Brackets{pro}, nil},
		{[]string{"herald", "guardian", "crusader", "archon",
			"legend", "ancient", "divine", "immortal", "pro"},
			Brackets(model.AllBrackets), nil,
		},
		{[]string{"divine", "immortal"}, Brackets{divine, immortal}, nil},
		// Duplicate inputs
		{[]string{"divine", "d", "7"}, Brackets{divine}, nil},
		{[]string{"immortal", "8", "7"}, Brackets{immortal, divine}, nil},
		{[]string{"ancient", "1", "2", "6"}, Brackets{ancient, herald, guardian}, nil},
		// Mixed inputs (valid and invalid)
		{[]string{"divine", "skilled", "herald", "bad"}, Brackets{divine, herald}, nil},
		{[]string{"herald", "guardian", "crusader", "archon",
			"legend", "ancient", "divine", "immortal", "pro", "bad", "1234"},
			Brackets(model.AllBrackets), nil,
		},
		// Invalid inputs that return errors
		{[]string{"20"}, Brackets{}, errNoValidBrackets},
		{[]string{"Skilled"}, Brackets{}, errNoValidBrackets},
		{[]string{"Bad"}, Brackets{}, errNoValidBrackets},
	}

	for _, in := range iw {
		got, err := parseBrackets(in.input)

		if in.err == nil && err != nil {
			t.Errorf("parseBrackets(%v)=%v failed, error: %v", in.input, in.want, err)
		} else if in.err != nil {
			if err == nil || in.err.Error() != err.Error() {
				t.Errorf("parseBrackets(%v) failed, expected error %v, got %v", in.input, in.err, err)
			}
		}

		var reason string
		if len(in.want) == len(got) {
			for i, b := range got {
				if in.want[i] != b {
					reason = fmt.Sprintf(
						"in.input[%d] != got[%d]. Got '%s', want '%s'.",
						i, i, b, in.want[i],
					)
					break
				}
			}
		} else {
			reason = fmt.Sprintf("Mismatched length between want and got for %v. Got %v, want %v", in.input, got, in.want)
		}

		if reason != "" {
			t.Errorf("parseBrackets(%v)=%v failed, reason: %s", in.input, in.want, reason)
		}
	}
}

var (
	mainstat   = model.LayoutMainStat
	single     = model.LayoutSingle
	attacktype = model.LayoutAttackType
	role       = model.LayoutRole
	legs       = model.LayoutLegs
	modify     = model.LayoutModify
)

func TestParseLayout(t *testing.T) {
	iw := []struct {
		input string
		want  *model.Layout
		err   error
	}{
		// Valid input
		{"mainstat", mainstat, nil},
		{"ms", mainstat, nil},
		{"stat", mainstat, nil},
		{"single", single, nil},
		{"s", single, nil},
		{"attack", attacktype, nil},
		{"a", attacktype, nil},
		{"role", role, nil},
		{"r", role, nil},
		{"legs", legs, nil},
		{"l", legs, nil},
		{"modify", modify, nil},
		{"m", modify, nil},
		{"none", modify, nil},
		{"n", modify, nil},
		// Invalid input
		{"cool", nil, fmt.Errorf(invalidLayout, "cool")},
	}
	for _, in := range iw {
		got, err := parseLayout(&in.input)

		if in.err == nil && err != nil {
			t.Errorf("parseLayout(%v) failed, error: %v", in.input, err)
		} else if in.err != nil {
			if err == nil || in.err.Error() != err.Error() {
				t.Errorf("parseLayout(%v) failed, expected error %v, got %v", in.input, in.err, err)
			}
		}

		if got != in.want {
			t.Errorf("parseLayout(%s) failed, got %v, want %v", in.input, got, in.want)
		}
	}
}
