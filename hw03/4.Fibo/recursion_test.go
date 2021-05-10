package fibo_test

import (
	"os"
	"path/filepath"
	"testing"

	fibo "github.com/balabanovds/otus-algo/hw03/4.Fibo"
	"github.com/balabanovds/otus-algo/tests"
)

type recursionFiboTest struct {
	n int
}

func (r *recursionFiboTest) Name() string {
	return "RecursionFibo"
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

func (r *recursionFiboTest) Path() string {
	cwd, _ := os.Getwd()
	return filepath.Join(cwd, "recursion_data")
}

func (r *recursionFiboTest) N() int {
	return r.n
}

func TestRecursionFibo(t *testing.T) {
	cwd, err := os.Getwd()
	tests.FatalOnErr(t, err)
	tests.RunTests(t, &recursionFiboTest{}, filepath.Join(cwd, "recursion_data"))
}
