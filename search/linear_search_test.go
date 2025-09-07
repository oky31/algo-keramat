package search

import "testing"

func TestLinearSearchInt(t *testing.T) {

	tests := []struct {
		haystack []int
		needle   int
		want     bool
	}{
		{[]int{10, 50, 1, 1, 5, 20, 122, 7, 8}, 5, true},
		{[]int{10, 50, 1, 1, 5, 20, 122, 7, 8}, 122, true},
		{[]int{10, 50, 1, 1, 5, 20, 122, 7, 8}, 99, false},
		{[]int{}, 1, false},
		{[]int{1}, 1, true},
		{[]int{2}, 1, false},
	}

	for _, tt := range tests {
		got := LinearSearchInt(tt.haystack, tt.needle)
		if got != tt.want {
			t.Errorf("LinearSearchInt(%v, %d) = %v; want %v", tt.haystack, tt.needle, got, tt.want)
		}
	}
}
