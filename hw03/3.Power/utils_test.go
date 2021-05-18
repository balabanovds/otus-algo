package pow_test

import (
	"strconv"
)

func parseIn(data []string) (float64, int, error) {
	x, err := strconv.ParseFloat(data[0], 64)
	if err != nil {
		return 0, 0, err
	}

	n, err := strconv.Atoi(data[1])
	if err != nil {
		return 0, 0, err
	}

	return x, n, nil
}
