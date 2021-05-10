package fibo

import "math/big"

func IterationFibo(n int) *big.Int {
	if n <= 1 {
		return new(big.Int).SetInt64(int64(n))
	}

	f0 := new(big.Int).SetInt64(0)
	f1 := new(big.Int).SetInt64(1)
	f2 := new(big.Int)

	for n > 1 {
		f2.Add(f0, f1)
		f0.Set(f1)
		f1.Set(f2)
		n--
	}

	return f2
}
