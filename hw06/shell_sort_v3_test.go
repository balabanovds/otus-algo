package sort_test

import sort "github.com/balabanovds/otus-algo/hw06"

type shellSortV3RandomTest struct {
	baseTestRandom
}

func newShellSortV3RandomTest(amount int) *shellSortV3RandomTest {
	s := &shellSortV3RandomTest{}
	s.setAmount(amount)
	return s
}

func (b shellSortV3RandomTest) Name() string {
	return "6.1-ShellSort_O(N^(4/3))_random"
}

func (b *shellSortV3RandomTest) Run(data []string) (string, error) {
	return b.run(data, sort.ShellSortV3)
}

type shellSortV3DigitsTest struct {
	baseTestDigits
}

func newShellSortV3DigitsTest(amount int) *shellSortV3DigitsTest {
	s := &shellSortV3DigitsTest{}
	s.setAmount(amount)
	return s
}

func (b shellSortV3DigitsTest) Name() string {
	return "6.2-ShellSort_O(N^(4/3))_digits"
}

func (b *shellSortV3DigitsTest) Run(data []string) (string, error) {
	return b.run(data, sort.ShellSortV3)
}

type shellSortV3SortedTest struct {
	baseTestSorted
}

func newShellSortV3SortedTest(amount int) *shellSortV3SortedTest {
	s := &shellSortV3SortedTest{}
	s.setAmount(amount)
	return s
}

func (b shellSortV3SortedTest) Name() string {
	return "6.3-ShellSort_O(N^(4/3))_sorted"
}

func (b *shellSortV3SortedTest) Run(data []string) (string, error) {
	return b.run(data, sort.ShellSortV3)
}

type shellSortV3ReverseTest struct {
	baseTestReverse
}

func newShellSortV3ReverseTest(amount int) *shellSortV3ReverseTest {
	s := &shellSortV3ReverseTest{}
	s.setAmount(amount)
	return s
}

func (b shellSortV3ReverseTest) Name() string {
	return "6.4-ShellSort_O(N^(4/3))_revers"
}

func (b *shellSortV3ReverseTest) Run(data []string) (string, error) {
	return b.run(data, sort.ShellSortV3)
}
