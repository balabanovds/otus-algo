package primes_test

import (
	"os"
	"path/filepath"
	"strconv"

	"github.com/balabanovds/otus-algo/tests"
)

type primesBase struct {
	n int
}

func (s *primesBase) Name() string {
	return "SimpleBruteForce"
}

func (s *primesBase) Run(_ []string) (string, error) {
	panic("implement me!")
}

func (s *primesBase) Cmp() tests.CmpFunc {
	return func(a, b string) bool {
		return a == b
	}
}

func (s *primesBase) Path() string {
	cwd, _ := os.Getwd()
	return filepath.Join(cwd, "bf_data")
}

func (s *primesBase) N() int {
	return s.n
}

func (s *primesBase) run(data []string, f func(int) int) (string, error) {
	n, err := strconv.Atoi(data[0])
	if err != nil {
		return "", err
	}

	res := f(n)
	s.n = n

	return strconv.Itoa(res), nil
}
