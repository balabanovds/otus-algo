package primes_test

import (
	"testing"

	"github.com/balabanovds/otus-algo/tests"
)

func TestPrimes(t *testing.T) {
	s := tests.NewSuite(t, "Primes",
		newSimpleTest(8),
		newImmidTest(9),
		newHalfPrimesTest(9),
		newSqrtPrimesTest(11),
		newOddPrimesTest(11),
		newMemPrimesTest(12),
	)

	s.Run().ReportFile("_report.txt")
}
