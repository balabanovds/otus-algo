package fibo_test

import (
	"os"
	"path/filepath"
	"strconv"

	fibo "github.com/balabanovds/otus-algo/hw03/4.Fibo"
	"github.com/balabanovds/otus-algo/tests"
)

type matrixFiboTest struct {
	n int
}

func (m *matrixFiboTest) Name() string {
	return "MatrixFibo"
}

func (m *matrixFiboTest) Run(data []string) (string, error) {
	n, err := strconv.Atoi(data[0])
	if err != nil {
		return "", err
	}

	res, err := fibo.MatrixFibo(n)
	if err != nil {
		return "", err
	}

	m.n = n

	return res.String(), nil
}

func (m *matrixFiboTest) Cmp() tests.CmpFunc {
	return func(a, b string) bool {
		return a == b
	}
}

func (m *matrixFiboTest) Path() string {
	cwd, _ := os.Getwd()
	return filepath.Join(cwd, "matrix_data")
}

func (m *matrixFiboTest) N() int {
	return m.n
}
