package cli

import "testing"

func TestParse(t *testing.T) {
	iw := struct {
		input []string
		want  error
	}{
		[]string{"-l", "mainstat"}, nil,
		[]string{"-l", "mainstat", "immortal"}, nil,
		[]string{"-l", "mainstat"}, nil,
		[]string{"-l", "mainstat"}, nil,
		[]string{"-l", "mainstat"}, nil,
		[]string{"-l", "mainstat"}, nil,
	}
}
