package pow_test

import (
	"fmt"
	"testing"

	pow "github.com/balabanovds/otus-algo/hw03/3.Power"
)

func TestPow(t *testing.T) {
	tests := []struct {
		n    int
		want float64
	}{
		{
			n:    2,
			want: 2.25,
		},
		{
			n:    5,
			want: 2.48832,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("n=%d;want=%f", tt.n, tt.want), func(t *testing.T) {
			got := pow.BinaryPow(1, tt.n)
			if got != tt.want {
				t.Fatalf("want: %f, got: %f", tt.want, got)
			}
		})
	}
}
