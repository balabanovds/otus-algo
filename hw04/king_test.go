package bitboard_test

import (
	"testing"

	bitboard "github.com/balabanovds/otus-algo/hw04"
	"github.com/balabanovds/otus-algo/tests"
)

type kingTest struct {
}

func (k kingTest) Name() string {
	return "KingTest"
}

func (k *kingTest) Run(data []string) (string, error) {
	return run(data, bitboard.King)
}

func (k kingTest) Cmp() tests.CmpFunc {
	return cmp()
}

func TestKing(t *testing.T) {
	tests.RunTests(t, &kingTest{}, path(t, "1.Bitboard_King"))
}
