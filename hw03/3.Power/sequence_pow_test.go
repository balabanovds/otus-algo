package pow_test

import (
	pow "github.com/balabanovds/otus-algo/hw03/3.Power"
)

type seqPowTest struct {
	baseTest
}

func newSeqTest(amount int) *seqPowTest {
	t := &seqPowTest{}
	t.setAmount(amount)
	return t
}

func (s *seqPowTest) Name() string {
	return "01-SequencePower"
}

func (s *seqPowTest) Run(data []string) (string, error) {
	return s.run(data, pow.SeqPow)
}
