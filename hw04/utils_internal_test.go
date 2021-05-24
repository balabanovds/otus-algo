package bitboard

import (
	"strconv"
	"testing"
)

func Test_coords(t *testing.T) {
	tests := []struct {
		pos int
		x   int
		y   int
	}{
		{pos: 9, x: 2, y: 2},
		{pos: 47, x: 8, y: 6},
		{pos: 63, x: 8, y: 8},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(strconv.Itoa(tt.pos), func(t *testing.T) {
			gotX, gotY := coords(tt.pos)
			if gotX != tt.x && gotY != tt.y {
				t.Fatalf("want: x=%d, y=%d, got: x=%d, y=%d", tt.x, tt.y, gotX, gotY)
			}
		})
	}
}
