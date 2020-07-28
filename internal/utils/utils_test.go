package utils

import (
	"fmt"
	"math"
	"testing"
)

func TestMaxInt(t *testing.T) {
	iw := []struct {
		input []int
		want1 int
		want2 error
	}{
		{[]int{1, 2, 3, 4, 5}, 5, nil},
		{[]int{5, -5}, 5, nil},
		{[]int{0, -1}, 0, nil},
		{[]int{1, -1}, 1, nil},
		{[]int{0}, 0, nil},
		{[]int{1}, 1, nil},
		{[]int{int(math.MaxInt32)}, int(math.MaxInt32), nil},
	}
	for _, in := range iw {
		if got, err := MaxInt(in.input); got != in.want1 || err != in.want2 {
			fmt.Printf(
				"MaxInt(%d), got (%d, %v), want (%d, %v)\n",
				in.input, got, err, in.want1, in.want2)
		}
	}
}
