package sort_test

import sort "github.com/balabanovds/otus-algo/hw06"

type selectionSortRandomTest struct {
	baseTestRandom
}

func newSelectionSortRandomTest(amount int) *selectionSortRandomTest {
	s := &selectionSortRandomTest{}
	s.setAmount(amount)
	return s
}

func (b selectionSortRandomTest) Name() string {
	return "2.1-SelectionSort_random"
}

func (b *selectionSortRandomTest) Run(data []string) (string, error) {
	return b.run(data, sort.SelectionSort)
}

type selectionSortDigitsTest struct {
	baseTestDigits
}

func newSelectionSortDigitsTest(amount int) *selectionSortDigitsTest {
	s := &selectionSortDigitsTest{}
	s.setAmount(amount)
	return s
}

func (b selectionSortDigitsTest) Name() string {
	return "2.2-SelectionSort_digits"
}

func (b *selectionSortDigitsTest) Run(data []string) (string, error) {
	return b.run(data, sort.SelectionSort)
}

type selectionSortSortedTest struct {
	baseTestSorted
}

func newSelectionSortSortedTest(amount int) *selectionSortSortedTest {
	s := &selectionSortSortedTest{}
	s.setAmount(amount)
	return s
}

func (b selectionSortSortedTest) Name() string {
	return "2.3-SelectionSort_sorted"
}

func (b *selectionSortSortedTest) Run(data []string) (string, error) {
	return b.run(data, sort.SelectionSort)
}

type selectionSortReverseTest struct {
	baseTestReverse
}

func newSelectionSortReverseTest(amount int) *selectionSortReverseTest {
	s := &selectionSortReverseTest{}
	s.setAmount(amount)
	return s
}

func (b selectionSortReverseTest) Name() string {
	return "2.4-SelectionSort_revers"
}

func (b *selectionSortReverseTest) Run(data []string) (string, error) {
	return b.run(data, sort.SelectionSort)
}
