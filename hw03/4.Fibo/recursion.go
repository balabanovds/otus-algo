package fibo

import "math/big"

func RecursionFibo(n int) *big.Int {
	i := new(big.Int)
	if n <= 1 {
		return i.SetInt64(int64(n))
	}

	return i.Add(RecursionFibo(n-1), RecursionFibo(n-2))
}
