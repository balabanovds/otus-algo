package tests

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
	"testing"
	"text/tabwriter"
	"time"
)

type SuiteTestRunner interface {
	TestRunner
	// Path where your test data is
	Path() string
	N() int
}

type Suite struct {
	t       *testing.T
	name    string
	runners []SuiteTestRunner
	results map[string][]testResult
	ns      map[int]struct{}
}

type testResult struct {
	n        int
	result   string
	duration time.Duration
}

func NewSuite(t *testing.T, name string, runners ...SuiteTestRunner) *Suite {
	return &Suite{
		t:       t,
		name:    name,
		runners: runners,
		results: make(map[string][]testResult),
		ns:      make(map[int]struct{}),
	}
}

func (s *Suite) AddRunner(runner SuiteTestRunner) *Suite {
	s.runners = append(s.runners, runner)
	return s
}

// Run all TestRunners regarding to files *.in, *.out in dir.
func (s *Suite) Run() *Suite {
	t := s.t

	for _, runner := range s.runners {
		inFiles, outFiles := findFilesInDir(t, runner.Path())
		for i := range inFiles {
			i := i
			runner := runner
			t.Run(runner.Name()+"-"+strconv.Itoa(i), func(t *testing.T) {
				got, dur := runTest(t, runner, inFiles[i], outFiles[i])
				s.results[runner.Name()] = append(s.results[runner.Name()], testResult{
					n:        runner.N(),
					result:   got,
					duration: dur,
				})
				s.ns[runner.N()] = struct{}{}
			})
		}
	}

	return s
}

func (s *Suite) ReportStdout() {
	s.report(os.Stdout)
}

func (s *Suite) ReportFile(filename string) {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0o755)
	FatalOnErr(s.t, err)
	defer func() { _ = f.Close() }()

	s.report(f)
}

func (s *Suite) report(out io.Writer) {
	keys := make([]int, 0, len(s.results))
	for n := range s.ns {
		keys = append(keys, n)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	twr := tabwriter.NewWriter(out, 5, 3, 2, ' ', tabwriter.AlignRight)
	row := []string{"N:"}
	for _, k := range keys {
		row = append(row, strconv.Itoa(k))
	}
	_, err := fmt.Fprintln(twr, strings.Join(row, "\t")+"\t")
	FatalOnErr(s.t, err)

	for name, testResults := range s.results {
		row := []string{name}
		for _, res := range testResults {
			row = append(row, res.duration.String())
		}

		tail := strings.Repeat("\tX", len(keys)-len(testResults)) + "\t"
		_, err = fmt.Fprintln(twr, strings.Join(row, "\t")+tail)
		FatalOnErr(s.t, err)
	}
	FatalOnErr(s.t, twr.Flush())
}
