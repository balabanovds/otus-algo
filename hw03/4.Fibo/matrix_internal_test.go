package fibo

import (
	"math/big"
	"reflect"
	"testing"
)

func TestMult(t *testing.T) {
	a := newMatrix(2, 2, big.NewInt(0), big.NewInt(1), big.NewInt(1), big.NewInt(1))
	b := newMatrix(2, 2, big.NewInt(1), big.NewInt(2), big.NewInt(2), big.NewInt(3))
	want := newMatrix(2, 2, big.NewInt(2), big.NewInt(3), big.NewInt(3), big.NewInt(5))

	got, err := matrixMult(a, b)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want=%v got=%v", want, got)
	}
}

func TestMult2(t *testing.T) {
	a := newMatrix(2, 1, big.NewInt(1), big.NewInt(1))
	b := newMatrix(2, 2, big.NewInt(0), big.NewInt(1), big.NewInt(1), big.NewInt(1))
	want := newMatrix(2, 1, big.NewInt(1), big.NewInt(2))

	got, err := matrixMult(a, b)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want=%v got=%v", want, got)
	}
}

func TestPow(t *testing.T) {
	a := newMatrix(2, 2, big.NewInt(0), big.NewInt(1), big.NewInt(1), big.NewInt(1))
	want := newMatrix(2, 2, big.NewInt(2), big.NewInt(3), big.NewInt(3), big.NewInt(5))

	got, err := matrixPow(a, 4)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want=%v got=%v", want, got)
	}
}
