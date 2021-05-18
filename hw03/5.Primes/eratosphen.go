package primes

func Eratosphen(n int) int {
	divs := make([]bool, n+1)
	cntr := 0

	for i := 2; i <= n; i++ {
		if !divs[i] {
			cntr++
			for j := i * i; j <= n; j += i {
				divs[j] = true
			}
		}
	}

	return cntr
}

func EratosphenLinear(n int) int {
	var pr []int
	lp := make([]int, n+1)

	for i := 2; i <= n; i++ {
		if lp[i] == 0 {
			lp[i] = i
			pr = append(pr, i)
		}
		for _, p := range pr {
			if p <= lp[i] && p*i <= n {
				lp[p*i] = p
			}
		}
	}

	return len(pr)
}
