package pow_test

import (
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
	result, n, err := run(data, pow.MultiPow)
	if err != nil {
		return "", err
	}

	m.n = n

	return result, nil
}

func (m *multiPowTest) Cmp() tests.CmpFunc {
	return cmp()
}

func (m *multiPowTest) N() int {
	return m.n
}
