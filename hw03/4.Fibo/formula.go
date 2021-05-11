package fibo

import (
	"math"
	"math/big"
)

func FormulaFibo(n int) *big.Int {
	fi := 0.5 + math.Sqrt(5)/2

	f := new(big.Float).SetFloat64(fi)
	p := pow(f, n)

	result := new(big.Float).Quo(p, new(big.Float).SetFloat64(math.Sqrt(5)))
	result = new(big.Float).Add(result, new(big.Float).SetFloat64(0.5))

	i, _ := result.Int(nil)

	return i
}

func pow(x *big.Float, n int) *big.Float {
	result := new(big.Float).SetInt64(1)

	for n > 0 {
		if n&1 == 1 {
			result.Mul(result, x)
		}
		x.Mul(x, x)
		n >>= 1
	}

	return result
}
