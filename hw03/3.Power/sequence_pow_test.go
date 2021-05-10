package pow_test

import (
	"path/filepath"

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
	result, n, err := run(data, pow.SeqPow)
	if err != nil {
		return "", err
	}

	s.n = n

	return result, nil
}

func (s *seqPowTest) Cmp() tests.CmpFunc {
	return cmp()
}

func (b *seqPowTest) Path() string {
	return filepath.Dir(".")
}

func (s *seqPowTest) N() int {
	return s.n
}
