package primes_test

import (
	"testing"

	"github.com/balabanovds/otus-algo/tests"
)

func TestPrimes(t *testing.T) {
	s := tests.NewSuite(t, "Primes",
		&simpleBruteForce{},
		&immidiatePrimesTest{},
		&halfPrimesTest{},
	)

	s.Run().ReportFile("_report.txt")
}
