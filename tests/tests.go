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
	Run([]string) string
}

func RunTests(t *testing.T, runner TestRunner, path string) {
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

	fatalOnErr(t, err)

	for i := range inFiles {
		t.Run(runner.Name()+"-"+strconv.Itoa(i), func(t *testing.T) {
			runTest(t, runner, inFiles[i], outFiles[i])
		})
	}
}

func runTest(t *testing.T, runner TestRunner, inFile string, outFile string) {
	inRaw, err := ioutil.ReadFile(inFile)
	fatalOnErr(t, err)

	outRaw, err := ioutil.ReadFile(outFile)
	fatalOnErr(t, err)

	input := strings.Split(string(inRaw), "\n")
	want := strings.TrimSpace(string(outRaw))

	t0 := time.Now()
	got := runner.Run(input)
	t1 := time.Since(t0)
	t.Logf("execution duration: %s\n", t1)

	if want != got {
		t.Errorf("FATAL: want: %q, got: %q", want, got)
		t.Logf("inFile: %s, outFile: %s\n", inFile, outFile)
	}
}

func fatalOnErr(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}
