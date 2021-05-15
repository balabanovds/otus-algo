package primes_test

import (
	"strconv"
	"testing"

	primes "github.com/balabanovds/otus-algo/hw03/5.Primes"
)

type simpleBruteForce struct {
	primesBase
}

func (s *simpleBruteForce) Name() string {
	return "SimpleBruteForce"
}

func (s *simpleBruteForce) Run(data []string) (string, error) {
	return s.run(data, primes.SimpleBruteForce)
}

func TestSimple(t *testing.T) {
	cases := []struct {
		input int
		want  int
	}{
		{1, 0},
		{2, 1},
		{3, 2},
		{4, 2},
		{10, 4},
		{100, 25},
		{1000, 168},
		{10000, 1229},
	}

	for _, tt := range cases {
		tt := tt
		t.Run(strconv.Itoa(tt.input), func(t *testing.T) {
			got := primes.SimpleBruteForce(tt.input)
			if tt.want != got {
				t.Errorf("want=%d, got=%d", tt.want, got)
			}
		})
	}
}
