package pow_test

import (
	pow "github.com/balabanovds/otus-algo/hw03/3.Power"
)

type multiPowTest struct {
	baseTest
}

func newMultiTest(amount int) *multiPowTest {
	t := &multiPowTest{}
	t.setAmount(amount)
	return t
}

func (m *multiPowTest) Name() string {
	return "02-MultiPow"
}

func (m *multiPowTest) Run(data []string) (string, error) {
	return m.run(data, pow.MultiPow)
}
