package tests

import (
	"io/fs"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
	"time"
)

type TestRunner interface {
	Name() string
	Run([]string) (string, error)
}

type cmpFunc func(a, b string) bool

func RunTests(t *testing.T, runner TestRunner, path string, cmp cmpFunc) {
	var inFiles []string
	var outFiles []string

	err := filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) == ".in" {
			inFiles = append(inFiles, path)
		}
		if filepath.Ext(path) == ".out" {
			outFiles = append(outFiles, path)
		}
		return nil
	})

	FatalOnErr(t, err)

	for i := range inFiles {
		i := i
		t.Run(runner.Name()+"-"+strconv.Itoa(i), func(t *testing.T) {
			runTest(t, runner, inFiles[i], outFiles[i], cmp)
		})
	}
}

func runTest(t *testing.T, runner TestRunner, inFile string, outFile string, cmp cmpFunc) {
	inRaw, err := ioutil.ReadFile(inFile)
	FatalOnErr(t, err)

	outRaw, err := ioutil.ReadFile(outFile)
	FatalOnErr(t, err)

	input := sanitizeClr(strings.Split(string(inRaw), "\n"))
	want := strings.TrimSpace(string(outRaw))

	t0 := time.Now()
	got, err := runner.Run(input)
	FatalOnErr(t, err)
	t1 := time.Since(t0)
	t.Logf("execution duration: %s\n", t1)

	if !cmp(want, got) {
		t.Errorf("FATAL: want: %q, got: %q", want, got)
		t.Logf("inFile: %s, outFile: %s\n", inFile, outFile)
	}
}

func FatalOnErr(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}

func sanitizeClr(in []string) []string {
	res := make([]string, 0, len(in))
	for _, v := range in {
		res = append(res, strings.TrimSpace(v))
	}
	return res
}
