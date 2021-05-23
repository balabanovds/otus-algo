package sort

import pow "github.com/balabanovds/otus-algo/hw03/3.Power"

// ShellSortV1 uses native Shell sequence [N/2^k] O(N^2)
func ShellSortV1(arr []int) {
	var (
		n    = len(arr)
		gaps []int
		k    = 1
	)

	for {
		gap := n / int(pow.BinaryPow(2, k))
		if gap <= 0 {
			break
		}
		gaps = append(gaps, gap)
		k++
	}
	for _, gap := range gaps {
		sortGap(gap, arr)
	}
}

// ShellSortV2 uses Hibbard gap sequence [2^k-1] O(N^1.5)
func ShellSortV2(arr []int) {
	var (
		n    = len(arr)
		gaps []int
		k    = 1
	)

	for {
		gap := int(pow.BinaryPow(2, k)) - 1
		if gap > n-1 {
			break
		}
		gaps = append(gaps, gap)
		k++
	}

	for g := len(gaps) - 1; g >= 0; g-- {
		gap := gaps[g]
		sortGap(gap, arr)
	}
}

// ShellSortV3 uses Sedgewick gap sequence [4^k+3*2^(k-1)+1] O(N^(4/3))
func ShellSortV3(arr []int) {
	var (
		n    = len(arr)
		gaps = []int{1}
		k    = 1
	)

	for {
		gap := int(pow.BinaryPow(4, k) + 3*pow.BinaryPow(2, k-1) + 1)
		if gap > n-1 {
			break
		}
		gaps = append(gaps, gap)
		k++
	}

	for g := len(gaps) - 1; g >= 0; g-- {
		gap := gaps[g]
		sortGap(gap, arr)
	}
}

func sortGap(gap int, arr []int) {
	n := len(arr)
	for i := 0; i+gap < n; i++ {
		j := i + gap
		tmp := arr[j]
		for j-gap >= 0 && arr[j-gap] > tmp {
			arr[j] = arr[j-gap]
			j -= gap
		}
		arr[j] = tmp
	}
}
