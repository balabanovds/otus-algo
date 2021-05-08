package pow

func MultiPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	if n == 1 {
		return x
	}

	result := x

	i := 1
	for {
		result *= result
		i++

		if 1<<i > n {
			break
		}
	}

	fromPow := 1 << (i - 1)

	for j := fromPow; j < n; j++ {
		result *= x
	}

	return result
}
