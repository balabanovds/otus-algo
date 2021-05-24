package bitboard_test

import (
	"testing"

	bitboard "github.com/balabanovds/otus-algo/hw04"
	"github.com/balabanovds/otus-algo/tests"
)

type bishopTest struct{}

func (b bishopTest) Name() string {
	return "BishopTest"
}

func (b *bishopTest) Run(data []string) (string, error) {
	return run(data, bitboard.Bishop)
}

func (b bishopTest) Cmp() tests.CmpFunc {
	return cmp()
}

func TestBishop(t *testing.T) {
	tests.RunTests(t, &bishopTest{}, path(t, "4.Bitboard_Bishop"))
}
