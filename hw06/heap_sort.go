package sort

func HeapSort(arr []int) {
	size := len(arr)
	for root := (size - 1) / 2; root >= 0; root-- {
		maxToRoot(arr, root, size)
	}

	for i := size - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		maxToRoot(arr, 0, i)
	}
}

func maxToRoot(arr []int, root, size int) {
	l := 2*root + 1
	r := 2*root + 2
	x := root
	if l < size && arr[x] < arr[l] {
		x = l
	}
	if r < size && arr[x] < arr[r] {
		x = r
	}

	if x == root {
		return
	}

	arr[x], arr[root] = arr[root], arr[x]
	maxToRoot(arr, x, size)
}
