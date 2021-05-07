package pow

func SeqPow(x float64, n int) float64 {
	result := 1.0

	for n > 0 {
		result *= x
		n--
	}

	return result
}
