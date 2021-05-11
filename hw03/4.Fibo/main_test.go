package fibo_test

import (
	"testing"

	"github.com/balabanovds/otus-algo/tests"
)

func TestFibo(t *testing.T) {
	s := tests.NewSuite(t, "FiboTests", &recursionFiboTest{}, &iterationFiboTest{}, &formulaFiboTest{})
	s.Run().ReportFile("_report.txt")
}
