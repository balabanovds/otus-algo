package bitboard_test

import (
	"testing"

	bitboard "github.com/balabanovds/otus-algo/hw04"
	"github.com/balabanovds/otus-algo/tests"
)

type queenTest struct{}

func (q queenTest) Name() string {
	return "QueenTest"
}

func (q queenTest) Run(data []string) (string, error) {
	return run(data, bitboard.Queen)
}

func (q queenTest) Cmp() tests.CmpFunc {
	return cmp()
}

func TestQueen(t *testing.T) {
	tests.RunTests(t, &queenTest{}, path(t, "5.Bitboard_Queen"))
}
