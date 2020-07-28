package cli

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

//// Custom type for parsing multiple values
//type brackets []int
//
//func (b *brackets) String() string {
//	return fmt.Sprintf("%v", []int(*b))
//}
//
//func (b *brackets) Set(val string) error {
//	v, err := strconv.Atoi(val)
//	if err != nil {
//		return err
//	}
//	*b = append(*b, v)
//	return nil
//}

func Parse() (*UserConfig, error) {
	var path string
	var err error

	name := flag.String("n", defaultGridName, "Grid name")
	layout := flag.String("l", defaultLayout, "Grid layout")
	path = *(flag.String("p", "", "Path to Dota 2 userdata directory"))
	if path == "" {
		path, err = autodetectUserdataDir()
		if err != nil {
			return nil, err
		}
	}
	sortAsc := flag.Bool("s", false, "Sort ascending (low-high) [default: high-low]")
	flag.Parse()

	b := flag.Args()

	if b == nil { // NOTE: is this the correct comparison?
		b = defaultBrackets
	}

	// Only include valid brackets
	//for _, bracket := range b {
	//	// Everything is title-cased
	//	bracket = strings.Title(bracket)
	//	if _, ok := Brackets[bracket]; ok {
	//		brackets = append(brackets, bracket)
	//	} else {
	//		// Just warn for now
	//		fmt.Fprintf(os.Stderr, "Argument '%s' is not a valid skill bracket\n", bracket)
	//	}
	//}
	b, err := parseBrackets(b)
	if err != nil {
		return nil, err
	}

	// Check layout validity
	if _, ok := Layouts[*layout]; !ok {
		return nil, fmt.Errorf("invalid layout '%s'", *layout)
	}

	return NewUserConfig(
		brackets,
		*name,
		*layout,
		path,
		*sortAsc,
	), nil
}


func parseBrackets(brackets []string]) ([]string, error) {
	var validBrackets []string
	// Only include valid brackets
	for _, bracket := range b {
		// Everything is title-cased
		bracket = strings.Title(bracket)
		if _, ok := Brackets[bracket]; ok {
			validBrackets = append(validBrackets, bracket)
		} else {
			return nil, fmt.Errorf("parse: argument '%s' is not a valid skill bracket", bracket)
		}
	}
	return validBrackets
}