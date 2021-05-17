package fibo_test

import (
	"testing"

	"github.com/balabanovds/otus-algo/tests"
)

func TestFibo(t *testing.T) {
	s := tests.NewSuite(t, "FiboTests",
		newRecursionTest(6),
		newIterationTest(11),
		newFormulaTest(12),
		newMatrixTest(12),
	)
	s.Run().ReportFile("_report.txt")
}
