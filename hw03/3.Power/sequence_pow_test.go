package pow_test

import (
	"fmt"
	"math"
	"strconv"

	pow "github.com/balabanovds/otus-algo/hw03/3.Power"

	"github.com/balabanovds/otus-algo/tests"
)

type seqPowTest struct {
	n int
}

func (s *seqPowTest) Name() string {
	return "SequencePower"
}

func (s *seqPowTest) Run(data []string) (string, error) {
	x, n, err := parseIn(data)
	if err != nil {
		return "", err
	}

	s.n = n

	result := pow.SeqPow(x, n)

	return fmt.Sprintf("%.06f", result), nil
}

func (s *seqPowTest) Cmp() tests.CmpFunc {
	return func(a, b string) bool {
		af, _ := strconv.ParseFloat(a, 64)
		bf, _ := strconv.ParseFloat(b, 64)
		var div float64 = 100

		return math.Floor(af*div)/div == math.Floor(bf*div)/div
	}
}

func (s *seqPowTest) N() int {
	return s.n
}
