package primes_test

import (
	"strconv"
	"testing"

	primes "github.com/balabanovds/otus-algo/hw03/5.Primes"
)

type immidiatePrimesTest struct {
	primesBase
}

func (s *immidiatePrimesTest) Name() string {
	return "ImmidiateBruteForce"
}

func (s *immidiatePrimesTest) Run(data []string) (string, error) {
	return s.run(data, primes.ImmidiateBruteForce)
}

type halfPrimesTest struct {
	primesBase
}

func (s *halfPrimesTest) Name() string {
	return "HalfBruteForce"
}

func (s *halfPrimesTest) Run(data []string) (string, error) {
	return s.run(data, primes.HalfBruteForce)
}

func TestImmidiateBruteForce(t *testing.T) {
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
			got := primes.ImmidiateBruteForce(tt.input)
			if tt.want != got {
				t.Errorf("want=%d, got=%d", tt.want, got)
			}
		})
	}
}
