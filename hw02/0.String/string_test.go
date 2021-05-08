package string_test

import (
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	"github.com/balabanovds/otus-algo/tests"
)

type stringTest struct{}

func (*stringTest) Run(input []string) (string, error) {
	l := len(strings.TrimSpace(input[0]))
	return strconv.Itoa(l), nil
}

func (*stringTest) Name() string {
	return "StringLength"
}

func (*stringTest) Cmp() tests.CmpFunc {
	return func(a, b string) bool { return a == b }
}

func TestString(t *testing.T) {
	tests.RunTests(t, &stringTest{}, filepath.Dir("."))
}
