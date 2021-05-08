package pow_test

import (
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
	result, n, err := run(data, pow.BinaryPow)
	if err != nil {
		return "", err
	}

	b.n = n

	return result, nil
}

func (b *binaryPowTest) Cmp() tests.CmpFunc {
	return cmp()
}

func (b *binaryPowTest) N() int {
	return b.n
}
