package pow_test

import (
	"fmt"
	"math"
	"strconv"

	"github.com/balabanovds/otus-algo/tests"

	pow "github.com/balabanovds/otus-algo/hw03/3.Power"
)

type binaryPowTest struct {
	n int
}

func (b *binaryPowTest) Name() string {
	return "BinaryPow"
}

func (b *binaryPowTest) Run(data []string) (string, error) {
	x, n, err := parseIn(data)
	if err != nil {
		return "", err
	}

	b.n = n

	result := pow.BinaryPow(x, n)

	return fmt.Sprintf("%.06f", result), nil
}

func (b *binaryPowTest) Cmp() tests.CmpFunc {
	return func(a, b string) bool {
		af, _ := strconv.ParseFloat(a, 64)
		bf, _ := strconv.ParseFloat(b, 64)
		var div float64 = 100

		return math.Floor(af*div)/div == math.Floor(bf*div)/div
	}
}

func (b *binaryPowTest) N() int {
	return b.n
}
