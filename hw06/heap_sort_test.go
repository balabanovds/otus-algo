package sort_test

import sort "github.com/balabanovds/otus-algo/hw06"

type heapSortRandomTest struct {
	baseTestRandom
}

func newHeapSortRandomTest(amount int) *heapSortRandomTest {
	s := &heapSortRandomTest{}
	s.setAmount(amount)
	return s
}

func (b heapSortRandomTest) Name() string {
	return "7.1-HeapSort_random"
}

func (b *heapSortRandomTest) Run(data []string) (string, error) {
	return b.run(data, sort.HeapSort)
}

type heapSortDigitsTest struct {
	baseTestDigits
}

func newHeapSortDigitsTest(amount int) *heapSortDigitsTest {
	s := &heapSortDigitsTest{}
	s.setAmount(amount)
	return s
}

func (b heapSortDigitsTest) Name() string {
	return "7.2-HeapSort_digits"
}

func (b *heapSortDigitsTest) Run(data []string) (string, error) {
	return b.run(data, sort.HeapSort)
}

type heapSortSortedTest struct {
	baseTestSorted
}

func newHeapSortSortedTest(amount int) *heapSortSortedTest {
	s := &heapSortSortedTest{}
	s.setAmount(amount)
	return s
}

func (b heapSortSortedTest) Name() string {
	return "7.3-HeapSort_sorted"
}

func (b *heapSortSortedTest) Run(data []string) (string, error) {
	return b.run(data, sort.HeapSort)
}

type heapSortReverseTest struct {
	baseTestReverse
}

func newHeapSortReverseTest(amount int) *heapSortReverseTest {
	s := &heapSortReverseTest{}
	s.setAmount(amount)
	return s
}

func (b heapSortReverseTest) Name() string {
	return "7.4-HeapSort_revers"
}

func (b *heapSortReverseTest) Run(data []string) (string, error) {
	return b.run(data, sort.HeapSort)
}
