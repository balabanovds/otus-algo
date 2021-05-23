// +build !ci

package sort_test

import (
	"testing"

	"github.com/balabanovds/otus-algo/tests"
)

func TestSort(t *testing.T) {
	s := tests.NewSuite(t, "Sorting",
		newBubbleSortRandomTest(5),
		newBubbleSortDigitsTest(5),
		newBubbleSortSortedTest(5),
		newBubbleSortReverseTest(5),
		newSelectionSortRandomTest(5),
		newSelectionSortDigitsTest(5),
		newSelectionSortSortedTest(5),
		newSelectionSortReverseTest(5),
		newInsertionSortRandomTest(5),
		newInsertionSortDigitsTest(5),
		newInsertionSortSortedTest(5),
		newInsertionSortReverseTest(5),
		newShellSortV1RandomTest(7),
		newShellSortV1DigitsTest(7),
		newShellSortV1SortedTest(7),
		newShellSortV1ReverseTest(7),
		newShellSortV2RandomTest(7),
		newShellSortV2DigitsTest(7),
		newShellSortV2SortedTest(7),
		newShellSortV2ReverseTest(7),
		newShellSortV3RandomTest(7),
		newShellSortV3DigitsTest(7),
		newShellSortV3SortedTest(7),
		newShellSortV3ReverseTest(7),
		newHeapSortRandomTest(7),
		newHeapSortDigitsTest(7),
		newHeapSortSortedTest(7),
		newHeapSortReverseTest(7),
	)
	s.Run().ReportFile("_report.txt")
}
