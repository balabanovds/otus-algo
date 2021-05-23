package sort_test

import sort "github.com/balabanovds/otus-algo/hw06"

type shellSortV2RandomTest struct {
	baseTestRandom
}

func newShellSortV2RandomTest(amount int) *shellSortV2RandomTest {
	s := &shellSortV2RandomTest{}
	s.setAmount(amount)
	return s
}

func (b shellSortV2RandomTest) Name() string {
	return "5.1-ShellSort_O(N^1.5)_random"
}

func (b *shellSortV2RandomTest) Run(data []string) (string, error) {
	return b.run(data, sort.ShellSortV2)
}

type shellSortV2DigitsTest struct {
	baseTestDigits
}

func newShellSortV2DigitsTest(amount int) *shellSortV2DigitsTest {
	s := &shellSortV2DigitsTest{}
	s.setAmount(amount)
	return s
}

func (b shellSortV2DigitsTest) Name() string {
	return "5.2-ShellSort_O(N^1.5)_digits"
}

func (b *shellSortV2DigitsTest) Run(data []string) (string, error) {
	return b.run(data, sort.ShellSortV2)
}

type shellSortV2SortedTest struct {
	baseTestSorted
}

func newShellSortV2SortedTest(amount int) *shellSortV2SortedTest {
	s := &shellSortV2SortedTest{}
	s.setAmount(amount)
	return s
}

func (b shellSortV2SortedTest) Name() string {
	return "5.3-ShellSort_O(N^1.5)_sorted"
}

func (b *shellSortV2SortedTest) Run(data []string) (string, error) {
	return b.run(data, sort.ShellSortV2)
}

type shellSortV2ReverseTest struct {
	baseTestReverse
}

func newShellSortV2ReverseTest(amount int) *shellSortV2ReverseTest {
	s := &shellSortV2ReverseTest{}
	s.setAmount(amount)
	return s
}

func (b shellSortV2ReverseTest) Name() string {
	return "5.4-ShellSort_O(N^1.5)_revers"
}

func (b *shellSortV2ReverseTest) Run(data []string) (string, error) {
	return b.run(data, sort.ShellSortV2)
}
