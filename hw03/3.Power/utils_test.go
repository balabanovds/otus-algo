package pow_test

import (
	"fmt"
	"math"
	"strconv"

	"github.com/balabanovds/otus-algo/tests"
)

func parseIn(data []string) (float64, int, error) {
	x, err := strconv.ParseFloat(data[0], 64)
	if err != nil {
		return 0, 0, err
	}

	n, err := strconv.Atoi(data[1])
	if err != nil {
		return 0, 0, err
	}

	return x, n, nil
}

func cmp() tests.CmpFunc {
	return func(a, b string) bool {
		af, _ := strconv.ParseFloat(a, 64)
		bf, _ := strconv.ParseFloat(b, 64)
		var div float64 = 100

		return math.Floor(af*div)/div == math.Floor(bf*div)/div
	}
}

func run(data []string, runF func(x float64, n int) float64) (result string, n int, err error) {
	var x float64
	x, n, err = parseIn(data)
	if err != nil {
		return
	}

	res := runF(x, n)

	result = fmt.Sprintf("%.06f", res)
	return
}
