package count

import "testing"

func TestCountSymbols(t *testing.T) {
	cases := []struct {
		input string
		want  int
	}{
		{"Hello", 5},
		{"", 0},
		{"Привет", 6},
		{"😊", 1}, // Emoji counts as one symbol
	}
	for _, c := range cases {
		got := CountSymbols(c.input)
		if got != c.want {
			t.Errorf("SymbolCount(%q) == %d, want %d", c.input, got, c.want)
		}
	}
}
