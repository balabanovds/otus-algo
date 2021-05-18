package primes_test

import (
	"os"
	"path/filepath"
	"strconv"

	"github.com/balabanovds/otus-algo/tests"
)

type baseTest struct {
	n      int
	amount int
}

func (s *baseTest) Name() string {
	return "SimpleBruteForce"
}

func (s *baseTest) Run(_ []string) (string, error) {
	panic("implement me!")
}

func (s *baseTest) Cmp() tests.CmpFunc {
	return func(a, b string) bool {
		return a == b
	}
}

func (s *baseTest) Path() string {
	cwd, _ := os.Getwd()
	return filepath.Join(cwd, "test_data")
}

func (s *baseTest) N() int {
	return s.n
}

func (s *baseTest) Amount() int {
	return s.amount
}

func (s *baseTest) setAmount(amount int) {
	s.amount = amount
}

func (s *baseTest) run(data []string, f func(int) int) (string, error) {
	n, err := strconv.Atoi(data[0])
	if err != nil {
		return "", err
	}

	res := f(n)
	s.n = n

	return strconv.Itoa(res), nil
}
