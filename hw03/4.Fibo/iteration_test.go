package fibo_test

import (
	"os"
	"path/filepath"

	"github.com/balabanovds/otus-algo/tests"

	fibo "github.com/balabanovds/otus-algo/hw03/4.Fibo"
)

type iterationFiboTest struct {
	amount int
	n      int
}

func newIterationTest(amount int) *iterationFiboTest {
	return &iterationFiboTest{amount: amount}
}

func (i *iterationFiboTest) Name() string {
	return "02-IterationFibo"
}

func (i *iterationFiboTest) Run(data []string) (string, error) {
	result, n, err := run(data, fibo.IterationFibo)
	if err != nil {
		return "", err
	}

	i.n = n
	return result, nil
}

func (i *iterationFiboTest) Cmp() tests.CmpFunc {
	return func(a, b string) bool {
		return a == b
	}
}

func (i *iterationFiboTest) Amount() int {
	return i.amount
}

func (i *iterationFiboTest) Path() string {
	cwd, _ := os.Getwd()
	return filepath.Join(cwd, "test_data")
}

func (i *iterationFiboTest) N() int {
	return i.n
}
