package fibo_test

import (
	"os"
	"path/filepath"

	fibo "github.com/balabanovds/otus-algo/hw03/4.Fibo"
	"github.com/balabanovds/otus-algo/tests"
)

type recursionFiboTest struct {
	amount int
	n      int
}

func newRecursionTest(amount int) *recursionFiboTest {
	return &recursionFiboTest{amount: amount}
}

func (r *recursionFiboTest) Name() string {
	return "01-RecursionFibo"
}

func (r *recursionFiboTest) Run(data []string) (string, error) {
	result, n, err := run(data, fibo.RecursionFibo)
	if err != nil {
		return "", err
	}

	r.n = n
	return result, nil
}

func (r *recursionFiboTest) Cmp() tests.CmpFunc {
	return func(a, b string) bool {
		return a == b
	}
}

func (r *recursionFiboTest) Amount() int {
	return r.amount
}

func (r *recursionFiboTest) Path() string {
	cwd, _ := os.Getwd()
	return filepath.Join(cwd, "test_data")
}

func (r *recursionFiboTest) N() int {
	return r.n
}
