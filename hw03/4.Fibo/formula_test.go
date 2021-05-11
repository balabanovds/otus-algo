package fibo_test

import (
	"math/big"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/balabanovds/otus-algo/tests"

	fibo "github.com/balabanovds/otus-algo/hw03/4.Fibo"
)

type formulaFiboTest struct {
	n int
}

func (f *formulaFiboTest) Name() string {
	return "FormulaFibo"
}

func (f *formulaFiboTest) Run(data []string) (string, error) {
	result, n, err := run(data, fibo.FormulaFibo)
	if err != nil {
		return "", err
	}

	f.n = n
	return result, nil
}

// здесь примерное сравнение
func (f *formulaFiboTest) Cmp() tests.CmpFunc {
	return func(a, b string) bool {
		bigA, ok := new(big.Int).SetString(a, 10)
		if !ok {
			return false
		}
		bigB, ok := new(big.Int).SetString(b, 10)
		if !ok {
			return false
		}

		lA := len(bigA.Bytes())
		lB := len(bigB.Bytes())
		if (lA != lB) && (!reflect.DeepEqual(bigA.Bytes()[:lA/2], bigB.Bytes()[:lB/2])) {
			return false
		}
		return true
	}
}

func (f *formulaFiboTest) Path() string {
	cwd, _ := os.Getwd()
	return filepath.Join(cwd, "formula_data")
}

func (f *formulaFiboTest) N() int {
	return f.n
}

func TestFormulaFibo(t *testing.T) {
	want, _ := new(big.Int).SetString("354224848179261915075", 10)
	got := fibo.FormulaFibo(100)

	lWant := len(want.Bytes())
	lGot := len(got.Bytes())
	if (lWant != lGot) && (!reflect.DeepEqual(want.Bytes()[:lWant/2], got.Bytes()[:lGot/2])) {
		t.Fatalf("want=%d, got=%d", want, got)
	}
}
