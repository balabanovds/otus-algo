package bitboard_test

import (
	"testing"

	bitboard "github.com/balabanovds/otus-algo/hw04"
	"github.com/balabanovds/otus-algo/tests"
)

type knightTest struct{}

func (k knightTest) Name() string {
	return "KnightTest"
}

func (k *knightTest) Run(data []string) (string, error) {
	return run(data, bitboard.Knight)
}

func (k knightTest) Cmp() tests.CmpFunc {
	return cmp()
}

func TestKnight(t *testing.T) {
	tests.RunTests(t, &knightTest{}, path(t, "2.Bitboard_Knight"))
}
