package string_test

import (
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	"github.com/balabanovds/otus-algo/tests"
)

type stringTest struct{}

func (*stringTest) Run(input []string) string {
	l := len(strings.TrimSpace(input[0]))
	return strconv.Itoa(l)
}

func (*stringTest) Name() string {
	return "StringLength"
}

func TestString(t *testing.T) {
	tests.RunTests(t, &stringTest{}, filepath.Dir("."))
}
