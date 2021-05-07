package pow_test

import (
	"fmt"
	"math"
	"path/filepath"
	"strconv"
	"testing"

	pow "github.com/balabanovds/otus-algo/hw03/3.Power"

	"github.com/balabanovds/otus-algo/tests"
)

type seqPowTest struct{}

func (s *seqPowTest) Name() string {
	return "SequencePower"
}

func (s *seqPowTest) Run(data []string) (string, error) {
	x, n, err := parseIn(data)
	if err != nil {
		return "", err
	}

	result := pow.SeqPow(x, n)

	return fmt.Sprintf("%.011f", result), nil
}

func TestSeqPow(t *testing.T) {
	cmp := func(a, b string) bool {
		af, _ := strconv.ParseFloat(a, 64)
		bf, _ := strconv.ParseFloat(b, 64)
		var div float64 = 100

		return math.Floor(af*div)/div == math.Floor(bf*div)/div
	}
	tests.RunTests(t, &seqPowTest{}, filepath.Dir("."), cmp)
}
