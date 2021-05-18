package fibo

import (
	"errors"
	"math/big"
)

func MatrixFibo(n int) (*big.Int, error) {
	p := newMatrix(2, 2, big.NewInt(0), big.NewInt(1), big.NewInt(1), big.NewInt(1))
	a := newMatrix(2, 1, big.NewInt(0), big.NewInt(1))

	pn, err := matrixPow(p, n)
	if err != nil {
		return nil, err
	}

	res, err := matrixMult(a, pn)
	if err != nil {
		return nil, err
	}

	return res[0][0], nil
}

type matrix [][]*big.Int

// newMatrix is filled from left ro right, from top to bottom with provided data
// if data omitted, then returned zero matrix
func newMatrix(x, y int, data ...*big.Int) matrix {
	m := make(matrix, y)

	cntr := 0

	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			if cntr < len(data) {
				m[i] = append(m[i], data[cntr])
			} else {
				m[i] = append(m[i], new(big.Int))
			}
			cntr++
		}
	}

	return m
}

func matrixMult(a, b matrix) (matrix, error) {
	if len(a[0]) != len(b) {
		return nil, errors.New("can't multiply matrixes")
	}

	out := make(matrix, len(a))
	for i := 0; i < len(a); i++ {
		out[i] = make([]*big.Int, len(b[0]))
		for j := 0; j < len(b[0]); j++ {
			out[i][j] = big.NewInt(0)
			for k := 0; k < len(b); k++ {
				mlt := new(big.Int).Mul(a[i][k], b[k][j])
				out[i][j] = new(big.Int).Add(out[i][j], mlt)
			}
		}
	}

	return out, nil
}

func matrixPow(m matrix, n int) (matrix, error) {
	result := newMatrix(2, 2, big.NewInt(1), big.NewInt(0), big.NewInt(0), big.NewInt(1))

	for n > 0 {
		var err error
		if n&1 == 1 {
			result, err = matrixMult(result, m)
			if err != nil {
				return nil, err
			}
		}
		m, err = matrixMult(m, m)
		if err != nil {
			return nil, err
		}
		n >>= 1
	}
	return result, nil
}
