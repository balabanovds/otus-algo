package tests

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
	"time"
)

type TestRunner interface {
	// Name of runner
	Name() string
	// Run main runner of test
	Run([]string) (string, error)
	// Cmp must return function that compares received result with wanted
	Cmp() CmpFunc
}

type CmpFunc func(a, b string) bool

func RunTests(t *testing.T, runner TestRunner, path string) {
	inFiles, outFiles := findFilesInDir(t, path)

	for i := range inFiles {
		i := i
		t.Run(runner.Name()+"-"+strconv.Itoa(i), func(t *testing.T) {
			_, _, _ = runTest(t, runner, inFiles[i], outFiles[i])
		})
	}
}

// runTest returns result of runner execution and duration of execution
func runTest(t *testing.T, runner TestRunner, inFile string, outFile string) (string, time.Duration, error) {
	input, want := getDataFromFiles(t, inFile, outFile)

	t0 := time.Now()
	got, err := runner.Run(input)
	FatalOnErr(t, err)
	t1 := time.Since(t0)
	t.Logf("execution duration: %s\n", t1)

	if !runner.Cmp()(want, got) {
		errMsg := fmt.Sprintf("want: %q, got: %q", want, got)
		t.Errorf("FATAL: %s", errMsg)
		t.Logf("inFile: %s, outFile: %s\n", inFile, outFile)
		return "", t1, errors.New(errMsg)
	}

	return got, t1, nil
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

func findFilesInDir(t *testing.T, path string) (inFiles []string, outFiles []string) {
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
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
	return
}

func getDataFromFiles(t *testing.T, inFile, outFile string) ([]string, string) {
	inRaw, err := ioutil.ReadFile(inFile)
	FatalOnErr(t, err)

	outRaw, err := ioutil.ReadFile(outFile)
	FatalOnErr(t, err)

	input := sanitizeClr(strings.Split(string(inRaw), "\n"))
	want := strings.TrimSpace(string(outRaw))

	return input, want
}
