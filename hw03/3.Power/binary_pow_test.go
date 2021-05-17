package pow_test

import (
	pow "github.com/balabanovds/otus-algo/hw03/3.Power"
)

type binaryPowTest struct {
	baseTest
}

func newBinaryTest(amount int) *binaryPowTest {
	t := &binaryPowTest{}
	t.setAmount(amount)
	return t
}

func (b *binaryPowTest) Name() string {
	return "03-BinaryPow"
}

func (b *binaryPowTest) Run(data []string) (string, error) {
	return b.run(data, pow.BinaryPow)
}
