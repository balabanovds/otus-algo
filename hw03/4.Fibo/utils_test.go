package fibo_test

import (
	"math/big"
	"strconv"
)

func run(data []string, runF func(n int) *big.Int) (result string, n int, err error) {
	n, err = strconv.Atoi(data[0])
	if err != nil {
		return
	}

	res := runF(n)

	result = res.String()
	return
}
