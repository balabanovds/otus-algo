package fibo_test

import (
	"os"
	"path/filepath"
	"strconv"

	fibo "github.com/balabanovds/otus-algo/hw03/4.Fibo"
	"github.com/balabanovds/otus-algo/tests"
)

type matrixFiboTest struct {
	amount int
	n      int
}

func newMatrixTest(amount int) *matrixFiboTest {
	return &matrixFiboTest{amount: amount}
}

func (m *matrixFiboTest) Name() string {
	return "04-MatrixFibo"
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

func (m *matrixFiboTest) Amount() int {
	return m.amount
}

func (m *matrixFiboTest) Path() string {
	cwd, _ := os.Getwd()
	return filepath.Join(cwd, "test_data")
}

func (m *matrixFiboTest) N() int {
	return m.n
}
