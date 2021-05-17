package pow_test

import (
	"testing"

	"github.com/balabanovds/otus-algo/tests"
)

func TestPow(t *testing.T) {
	s := tests.NewSuite(t, "Power",
		newSeqTest(9),
		newMultiTest(9),
		newBinaryTest(9),
	)
	s.Run().ReportFile("_report.txt")
}
