package pow

func BinaryPow(x float64, n int) float64 {
	result := 1.0

	for n > 0 {
		if n&1 == 1 {
			result *= x
		}
		x *= x
		n >>= 1
	}
	return result
}
