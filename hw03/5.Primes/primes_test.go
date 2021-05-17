package primes_test

import (
	primes "github.com/balabanovds/otus-algo/hw03/5.Primes"
)

type simpleBruteForce struct {
	amount int
	baseTest
}

func newSimpleTest(amount int) *simpleBruteForce {
	s := simpleBruteForce{amount: amount}
	s.setAmount(amount)
	return &s
}

func (s *simpleBruteForce) Name() string {
	return "01-SimpleBruteForce"
}

func (s *simpleBruteForce) Run(data []string) (string, error) {
	return s.run(data, primes.SimpleBruteForce)
}

type immidiatePrimesTest struct {
	baseTest
}

func newImmidTest(amount int) *immidiatePrimesTest {
	s := immidiatePrimesTest{}
	s.setAmount(amount)
	return &s
}

func (s *immidiatePrimesTest) Name() string {
	return "02-ImmidiateBruteForce"
}

func (s *immidiatePrimesTest) Run(data []string) (string, error) {
	return s.run(data, primes.ImmidiateBruteForce)
}

type halfPrimesTest struct {
	baseTest
}

func newHalfPrimesTest(amount int) *halfPrimesTest {
	s := halfPrimesTest{}
	s.setAmount(amount)
	return &s
}

func (s *halfPrimesTest) Name() string {
	return "03-HalfBruteForce"
}

func (s *halfPrimesTest) Run(data []string) (string, error) {
	return s.run(data, primes.HalfBruteForce)
}

type sqrtPrimesTest struct {
	baseTest
}

func newSqrtPrimesTest(amount int) *sqrtPrimesTest {
	s := sqrtPrimesTest{}
	s.setAmount(amount)
	return &s
}

func (s *sqrtPrimesTest) Name() string {
	return "04-SqrtBruteForce"
}

func (s *sqrtPrimesTest) Run(data []string) (string, error) {
	return s.run(data, primes.SqrtBruteForce)
}

type oddPrimesTest struct {
	baseTest
}

func newOddPrimesTest(amount int) *oddPrimesTest {
	s := oddPrimesTest{}
	s.setAmount(amount)
	return &s
}

func (s *oddPrimesTest) Name() string {
	return "05-OddBruteForce"
}

func (s *oddPrimesTest) Run(data []string) (string, error) {
	return s.run(data, primes.OddBruteForce)
}

type memPrimesTest struct {
	baseTest
}

func newMemPrimesTest(amount int) *memPrimesTest {
	s := memPrimesTest{}
	s.setAmount(amount)
	return &s
}

func (s *memPrimesTest) Name() string {
	return "06-MemoryPrimes"
}

func (s *memPrimesTest) Run(data []string) (string, error) {
	return s.run(data, primes.MemoryPrimes)
}
