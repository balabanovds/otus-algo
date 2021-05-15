package primes

func SimpleBruteForce(n int) int {
	isPrime := func(i int) bool {
		cntr := 0
		for j := 2; j <= i; j++ {
			if i%j == 0 {
				cntr++
			}
		}
		return cntr == 1
	}

	return iterateOver(n, isPrime)
}

func ImmidiateBruteForce(n int) int {
	isPrime := func(i int) bool {
		for j := 2; j < i; j++ {
			if i%j == 0 {
				return false
			}
		}
		return true
	}

	return iterateOver(n, isPrime)
}

func HalfBruteForce(n int) int {
	isPrime := func(i int) bool {
		for j := 2; j <= i/2; j++ {
			if i%j == 0 {
				return false
			}
		}
		return true
	}

	return iterateOver(n, isPrime)
}

func iterateOver(n int, isPrime func(i int) bool) int {
	res := 0
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			res++
		}
	}
	return res
}
