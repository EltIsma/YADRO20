package internal

import "testing"

func TestChecksortabilityTabliDriven(t *testing.T) {
	var tests = []struct {
		input [][]int32
		want  bool
	}{
		// the table itself
		{[][]int32{
			{1, 2},
			{2, 1},
		}, true},
		{[][]int32{
			{1, 0},
			{0, 1},
		}, true},
		{[][]int32{
			{0, 4, 0},
			{2, 0, 1},
			{1, 0, 2},
		}, true},
		{[][]int32{
			{2, 1},
			{0, 0},
		}, false},
		{[][]int32{
			{10, 20, 30},
			{1, 1, 1},
			{0, 0, 1},
		}, false},
		{[][]int32{
			{1, 2, 3},
			{3, 2, 1},
			{2, 3, 1},
		}, false},
	}
	for _, tt := range tests {
		t.Run("checkSortability", func(t *testing.T) {
			ans := Checksortability(tt.input)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
