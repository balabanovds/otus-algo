package sort_test

import sort "github.com/balabanovds/otus-algo/hw06"

type shellSortV1RandomTest struct {
	baseTestRandom
}

func newShellSortV1RandomTest(amount int) *shellSortV1RandomTest {
	s := &shellSortV1RandomTest{}
	s.setAmount(amount)
	return s
}

func (b shellSortV1RandomTest) Name() string {
	return "4.1-ShellSort_O(N^2)_random"
}

func (b *shellSortV1RandomTest) Run(data []string) (string, error) {
	return b.run(data, sort.ShellSortV1)
}

type shellSortV1DigitsTest struct {
	baseTestDigits
}

func newShellSortV1DigitsTest(amount int) *shellSortV1DigitsTest {
	s := &shellSortV1DigitsTest{}
	s.setAmount(amount)
	return s
}

func (b shellSortV1DigitsTest) Name() string {
	return "4.2-ShellSort_O(N^2)_digits"
}

func (b *shellSortV1DigitsTest) Run(data []string) (string, error) {
	return b.run(data, sort.ShellSortV1)
}

type shellSortV1SortedTest struct {
	baseTestSorted
}

func newShellSortV1SortedTest(amount int) *shellSortV1SortedTest {
	s := &shellSortV1SortedTest{}
	s.setAmount(amount)
	return s
}

func (b shellSortV1SortedTest) Name() string {
	return "4.3-ShellSort_O(N^2)_sorted"
}

func (b *shellSortV1SortedTest) Run(data []string) (string, error) {
	return b.run(data, sort.ShellSortV1)
}

type shellSortV1ReverseTest struct {
	baseTestReverse
}

func newShellSortV1ReverseTest(amount int) *shellSortV1ReverseTest {
	s := &shellSortV1ReverseTest{}
	s.setAmount(amount)
	return s
}

func (b shellSortV1ReverseTest) Name() string {
	return "4.4-ShellSort_O(N^2)_revers"
}

func (b *shellSortV1ReverseTest) Run(data []string) (string, error) {
	return b.run(data, sort.ShellSortV1)
}
