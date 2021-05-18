package primes_test

import primes "github.com/balabanovds/otus-algo/hw03/5.Primes"

type eratosphenTest struct {
	amount int
	baseTest
}

func newEratosphenTest(amount int) *eratosphenTest {
	s := eratosphenTest{amount: amount}
	s.setAmount(amount)
	return &s
}

func (s *eratosphenTest) Name() string {
	return "07-Eratosphen"
}

func (s *eratosphenTest) Run(data []string) (string, error) {
	return s.run(data, primes.Eratosphen)
}

type eratosphenTestLinear struct {
	amount int
	baseTest
}

func newEratosphenLinearTest(amount int) *eratosphenTestLinear {
	s := eratosphenTestLinear{amount: amount}
	s.setAmount(amount)
	return &s
}

func (s *eratosphenTestLinear) Name() string {
	return "07-EratosphenLinear"
}

func (s *eratosphenTestLinear) Run(data []string) (string, error) {
	return s.run(data, primes.EratosphenLinear)
}
