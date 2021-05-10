package pow_test

import (
	"testing"

	"github.com/balabanovds/otus-algo/tests"
)

func TestPow(t *testing.T) {
	s := tests.NewSuite(t, "Power", &seqPowTest{}, &multiPowTest{}, &binaryPowTest{})
	s.Run().ReportFile("_report.txt")
}
