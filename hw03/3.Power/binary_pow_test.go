package pow_test

import (
	"fmt"
	"math"
	"path/filepath"
	"strconv"
	"testing"

	"github.com/balabanovds/otus-algo/tests"

	pow "github.com/balabanovds/otus-algo/hw03/3.Power"
)

type binaryPowTest struct{}

func (b *binaryPowTest) Name() string {
	return "BinaryPow"
}

func (b *binaryPowTest) Run(data []string) (string, error) {
	x, n, err := parseIn(data)
	if err != nil {
		return "", err
	}

	result := pow.BinaryPow(x, n)

	return fmt.Sprintf("%.06f", result), nil
}

func TestBinaryPow(t *testing.T) {
	cmp := func(a, b string) bool {
		af, _ := strconv.ParseFloat(a, 64)
		bf, _ := strconv.ParseFloat(b, 64)
		var div float64 = 100

		return math.Floor(af*div)/div == math.Floor(bf*div)/div
	}
	tests.RunTests(t, &binaryPowTest{}, filepath.Dir("."), cmp)
}
