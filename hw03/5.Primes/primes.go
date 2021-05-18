package primes

import "math"

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

func SqrtBruteForce(n int) int {
	isPrime := func(i int) bool {
		to := int(math.Sqrt(float64(i)))
		for j := 2; j <= to; j++ {
			if i%j == 0 {
				return false
			}
		}
		return true
	}

	return iterateOver(n, isPrime)
}

func OddBruteForce(n int) int {
	isPrime := func(i int) bool {
		if i == 2 {
			return true
		}

		if i%2 == 0 {
			return false
		}

		to := int(math.Sqrt(float64(i)))
		for j := 3; j <= to; j += 2 {
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

func MemoryPrimes(n int) int {
	if n < 2 {
		return 0
	}
	primes := []int{2}

	isPrime := func(i int) bool {
		to := int(math.Sqrt(float64(i)))
		for j := 0; primes[j] <= to; j++ {
			if i%primes[j] == 0 {
				return false
			}
		}
		return true
	}

	for i := 3; i <= n; i += 2 {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}
	return len(primes)
}
