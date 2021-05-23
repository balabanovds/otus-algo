package sort_test

import sort "github.com/balabanovds/otus-algo/hw06"

type insertionSortRandomTest struct {
	baseTestRandom
}

func newInsertionSortRandomTest(amount int) *insertionSortRandomTest {
	s := &insertionSortRandomTest{}
	s.setAmount(amount)
	return s
}

func (b insertionSortRandomTest) Name() string {
	return "3.1-InsertionSort_random"
}

func (b *insertionSortRandomTest) Run(data []string) (string, error) {
	return b.run(data, sort.InsertionSort)
}

type insertionSortDigitsTest struct {
	baseTestDigits
}

func newInsertionSortDigitsTest(amount int) *insertionSortDigitsTest {
	s := &insertionSortDigitsTest{}
	s.setAmount(amount)
	return s
}

func (b insertionSortDigitsTest) Name() string {
	return "3.2-InsertionSort_digits"
}

func (b *insertionSortDigitsTest) Run(data []string) (string, error) {
	return b.run(data, sort.InsertionSort)
}

type insertionSortSortedTest struct {
	baseTestSorted
}

func newInsertionSortSortedTest(amount int) *insertionSortSortedTest {
	s := &insertionSortSortedTest{}
	s.setAmount(amount)
	return s
}

func (b insertionSortSortedTest) Name() string {
	return "3.3-InsertionSort_sorted"
}

func (b *insertionSortSortedTest) Run(data []string) (string, error) {
	return b.run(data, sort.InsertionSort)
}

type insertionSortReverseTest struct {
	baseTestReverse
}

func newInsertionSortReverseTest(amount int) *insertionSortReverseTest {
	s := &insertionSortReverseTest{}
	s.setAmount(amount)
	return s
}

func (b insertionSortReverseTest) Name() string {
	return "3.4-InsertionSort_revers"
}

func (b *insertionSortReverseTest) Run(data []string) (string, error) {
	return b.run(data, sort.InsertionSort)
}
