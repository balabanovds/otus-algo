package fibo_test

import (
	"os"
	"path/filepath"

	"github.com/balabanovds/otus-algo/tests"

	fibo "github.com/balabanovds/otus-algo/hw03/4.Fibo"
)

type iterationFiboTest struct {
	n int
}

func (i *iterationFiboTest) Name() string {
	return "IterationFibo"
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

func (i *iterationFiboTest) Path() string {
	cwd, _ := os.Getwd()
	return filepath.Join(cwd, "iteration_data")
}

func (i *iterationFiboTest) N() int {
	return i.n
}
