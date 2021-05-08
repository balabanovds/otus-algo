package pow_test

import (
	"fmt"
	"math"
	"strconv"

	pow "github.com/balabanovds/otus-algo/hw03/3.Power"
	"github.com/balabanovds/otus-algo/tests"
)

type multiPowTest struct {
	n int
}

func (m *multiPowTest) Name() string {
	return "MultiPow"
}

func (m *multiPowTest) Run(data []string) (string, error) {
	x, n, err := parseIn(data)
	if err != nil {
		return "", err
	}

	m.n = n

	result := pow.MultiPow(x, n)

	return fmt.Sprintf("%.06f", result), nil
}

func (m *multiPowTest) Cmp() tests.CmpFunc {
	return func(a, b string) bool {
		af, _ := strconv.ParseFloat(a, 64)
		bf, _ := strconv.ParseFloat(b, 64)
		var div float64 = 100

		return math.Floor(af*div)/div == math.Floor(bf*div)/div
	}
}

func (m *multiPowTest) N() int {
	return m.n
}
