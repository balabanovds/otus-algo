package pow_test

import (
	"fmt"
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
	x, err := strconv.ParseFloat(data[0], 64)
	if err != nil {
		return "", err
	}

	n, err := strconv.Atoi(data[1])
	if err != nil {
		return "", err
	}

	result := pow.BinaryPow(x, n)

	return fmt.Sprintf("%.011f", result), nil
}

func TestBinaryPow(t *testing.T) {
	tests.RunTests(t, &binaryPowTest{}, filepath.Dir("."))
}
