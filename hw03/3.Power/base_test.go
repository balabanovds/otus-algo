package pow_test

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strconv"

	"github.com/balabanovds/otus-algo/tests"
)

type baseTest struct {
	amount int
	n      int
}

func (b *baseTest) Name() string {
	panic("implement me")
}

func (b *baseTest) Run(_ []string) (string, error) {
	panic("implement me")
}

func (b *baseTest) run(data []string, f func(x float64, n int) float64) (string, error) {
	x, n, err := parseIn(data)
	if err != nil {
		return "", err
	}

	res := f(x, n)
	b.n = n

	return fmt.Sprintf("%.011f", res), nil
}

func (b *baseTest) Cmp() tests.CmpFunc {
	return func(a, b string) bool {
		af, _ := strconv.ParseFloat(a, 64)
		bf, _ := strconv.ParseFloat(b, 64)
		var div float64 = 100

		return math.Floor(af*div)/div == math.Floor(bf*div)/div
	}
}

func (b *baseTest) Path() string {
	cwd, _ := os.Getwd()
	return filepath.Join(cwd, "test_data")
}

func (b *baseTest) Amount() int {
	return b.amount
}

func (b *baseTest) setAmount(amount int) {
	b.amount = amount
}

func (b baseTest) N() int {
	return b.n
}
