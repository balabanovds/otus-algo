package sort_test

import (
	"errors"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"

	"github.com/balabanovds/otus-algo/tests"
)

type baseTest struct {
	n      int
	amount int
}

func (b baseTest) Name() string {
	panic("implement me")
}

func (b baseTest) Run(_ []string) (string, error) {
	panic("implement me")
}

func (b baseTest) Cmp() tests.CmpFunc {
	return func(a, b string) bool {
		_a := strings.Fields(a)
		_b := strings.Fields(b)
		return reflect.DeepEqual(_a, _b)
	}
}

func (b baseTest) Path() string {
	panic("implement me")
}

func (b baseTest) Amount() int {
	return b.amount
}

func (b baseTest) N() int {
	return b.n
}

func (b *baseTest) setAmount(amount int) {
	b.amount = amount
}

func (b *baseTest) run(data []string, fSort func([]int)) (string, error) {
	if len(data) < 2 {
		return "", errors.New("not valid data length")
	}
	n, err := strconv.Atoi(data[0])
	if err != nil {
		return "", err
	}
	b.n = n

	_arr := strings.Fields(data[1])
	arr := make([]int, 0, len(_arr))
	for _, a := range _arr {
		i, err := strconv.Atoi(a)
		if err != nil {
			return "", err
		}
		arr = append(arr, i)
	}

	fSort(arr)

	for i := range arr {
		_arr[i] = strconv.Itoa(arr[i])
	}

	return strings.Join(_arr, " "), nil
}

type baseTestRandom struct {
	baseTest
}

func (b baseTestRandom) Path() string {
	cwd, _ := os.Getwd()
	return filepath.Join(cwd, "test_data", "0.random")
}

type baseTestDigits struct {
	baseTest
}

func (b baseTestDigits) Path() string {
	cwd, _ := os.Getwd()
	return filepath.Join(cwd, "test_data", "1.digits")
}

type baseTestSorted struct {
	baseTest
}

func (b baseTestSorted) Path() string {
	cwd, _ := os.Getwd()
	return filepath.Join(cwd, "test_data", "2.sorted")
}

type baseTestReverse struct {
	baseTest
}

func (b baseTestReverse) Path() string {
	cwd, _ := os.Getwd()
	return filepath.Join(cwd, "test_data", "3.revers")
}
