package bitboard_test

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"testing"

	"github.com/balabanovds/otus-algo/tests"
)

func run(data []string, runF func(int) (int, uint64)) (string, error) {
	pos, err := strconv.Atoi(data[0])
	if err != nil {
		return "", err
	}

	cntr, bMoves := runF(pos)
	return fmt.Sprintf("%d\r\n%d", cntr, bMoves), nil
}

func cmp() tests.CmpFunc {
	return func(a, b string) bool {
		return a == b
	}
}

func path(t *testing.T, dir string) string {
	cwd, err := os.Getwd()
	tests.FatalOnErr(t, err)
	return filepath.Join(cwd, "test_data", dir)
}
