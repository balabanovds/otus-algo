package sort_test

import sort "github.com/balabanovds/otus-algo/hw06"

type bubbleSortRandomTest struct {
	baseTestRandom
}

func newBubbleSortRandomTest(amount int) *bubbleSortRandomTest {
	s := &bubbleSortRandomTest{}
	s.setAmount(amount)
	return s
}

func (b bubbleSortRandomTest) Name() string {
	return "1.1-BubbleSort_random"
}

func (b *bubbleSortRandomTest) Run(data []string) (string, error) {
	return b.run(data, sort.BubbleSort)
}

type bubbleSortDigitsTest struct {
	baseTestDigits
}

func newBubbleSortDigitsTest(amount int) *bubbleSortDigitsTest {
	s := &bubbleSortDigitsTest{}
	s.setAmount(amount)
	return s
}

func (b bubbleSortDigitsTest) Name() string {
	return "1.2-BubbleSort_digits"
}

func (b *bubbleSortDigitsTest) Run(data []string) (string, error) {
	return b.run(data, sort.BubbleSort)
}

type bubbleSortSortedTest struct {
	baseTestSorted
}

func newBubbleSortSortedTest(amount int) *bubbleSortSortedTest {
	s := &bubbleSortSortedTest{}
	s.setAmount(amount)
	return s
}

func (b bubbleSortSortedTest) Name() string {
	return "1.3-BubbleSort_sorted"
}

func (b *bubbleSortSortedTest) Run(data []string) (string, error) {
	return b.run(data, sort.BubbleSort)
}

type bubbleSortReverseTest struct {
	baseTestReverse
}

func newBubbleSortReverseTest(amount int) *bubbleSortReverseTest {
	s := &bubbleSortReverseTest{}
	s.setAmount(amount)
	return s
}

func (b bubbleSortReverseTest) Name() string {
	return "1.4-BubbleSort_revers"
}

func (b *bubbleSortReverseTest) Run(data []string) (string, error) {
	return b.run(data, sort.BubbleSort)
}
