package pow_test

import (
	"path/filepath"
	"testing"

	"github.com/balabanovds/otus-algo/tests"
)

func TestPow(t *testing.T) {
	s := tests.NewSuite(t, "Power", &seqPowTest{}, &binaryPowTest{})
	s.Run(filepath.Dir("."))
	s.ReportFile("power_report.txt")
}
