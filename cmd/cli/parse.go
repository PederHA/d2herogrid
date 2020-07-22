package cli

import (
	"flag"
	"fmt"
	"strconv"
)

type brackets []int

func (b *brackets) String() string {
	return fmt.Sprintf("%v", []int(*b))
}

func (b *brackets) Set(val string) error {
	v, err := strconv.Atoi(val)
	if err != nil {
		return err
	}
	*b = append(*b, v)
	return nil
}

func Parse() *UserConfig {
	var b brackets
	flag.Var(&b, "b", "Bracket")
	name := flag.String("n", defaultGridName, "Grid name")
	layout := flag.Int("l", defaultLayout, fmt.Sprintf("Grid layout (0-%d)", LayoutRole))
	path := flag.String("p", "", "Dota userdata directory path")
	sortAsc := flag.Bool("s", false, "Sort ascending (low-high)")
	flag.Parse()

	return NewUserConfig(
		[]int(b),
		*name,
		*layout,
		*path,
		*sortAsc,
	)
}
