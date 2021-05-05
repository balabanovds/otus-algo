package tickets_test

import (
	"path/filepath"
	"strconv"
	"testing"

	"github.com/balabanovds/otus-algo/tests"
)

type recursiveTest struct {
	digitsNum       int
	ticketsQuantity int64
}

func (r *recursiveTest) Name() string {
	return "TicketsRecursive"
}

func (r *recursiveTest) Run(input []string) (string, error) {
	var err error
	r.digitsNum, err = strconv.Atoi(input[0])
	if err != nil {
		return "", err
	}

	r.nextDigit(0, 0, 0)

	return strconv.Itoa(int(r.ticketsQuantity)), nil
}

func (r *recursiveTest) nextDigit(digit int, sum1 int, sum2 int) {
	if digit == r.digitsNum {
		if sum1 == sum2 {
			r.ticketsQuantity++
		}
		return
	}

	for a := 0; a < 10; a++ {
		for b := 0; b < 10; b++ {
			r.nextDigit(digit+1, sum1+a, sum2+b)
		}
	}
}

func TestLuckyTickets_recursive(t *testing.T) {
	tests.RunTests(t, &recursiveTest{}, filepath.Dir("."))
}
