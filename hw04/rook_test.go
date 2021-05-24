package bitboard_test

import (
	"testing"

	bitboard "github.com/balabanovds/otus-algo/hw04"
	"github.com/balabanovds/otus-algo/tests"
)

type rookTest struct{}

func (r rookTest) Name() string {
	return "RookTest"
}

func (r *rookTest) Run(data []string) (string, error) {
	return run(data, bitboard.Rook)
}

func (r rookTest) Cmp() tests.CmpFunc {
	return cmp()
}

func TestRook(t *testing.T) {
	tests.RunTests(t, &rookTest{}, path(t, "3.Bitboard_Rook"))
}
