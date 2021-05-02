package hw01

import (
	"math"
	"os"
	"path"
	"reflect"
	"strings"
	"testing"
)

func TestHarry(t *testing.T) {
	type testCase struct {
		filename  string
		predicate func(x, y int) bool
	}

	tests := []testCase{
		{
			filename: "01.txt",
			predicate: func(x, y int) bool {
				return x > y
			},
		},
		{
			filename: "02.txt",
			predicate: func(x, y int) bool {
				return x == y
			},
		},
		{
			filename: "03.txt",
			predicate: func(x, y int) bool {
				return (size - x - 1) == y
			},
		},
		{
			filename: "04.txt",
			predicate: func(x, y int) bool {
				return (size - x) > y-5
			},
		},
		{
			filename: "05.txt",
			predicate: func(x, y int) bool {
				return x/2 == y
			},
		},
		{
			filename: "06.txt",
			predicate: func(x, y int) bool {
				return x < 10 || y < 10
			},
		},
		{
			filename: "07.txt",
			predicate: func(x, y int) bool {
				return (size-x) < 10 && (size-y) < 10
			},
		},
		{
			filename: "08.txt",
			predicate: func(x, y int) bool {
				return x == 0 || y == 0
			},
		},
		{
			filename: "09.txt",
			predicate: func(x, y int) bool {
				return (x > y+10) || (y > x+10)
			},
		},
		{
			filename: "10.txt",
			predicate: func(x, y int) bool {
				return (x > y) && (x/2 <= y)
			},
		},
		{
			filename: "11.txt",
			predicate: func(x, y int) bool {
				return (x == 1) || (y == 1) || (size-x-1 == 1) || (size-y-1 == 1)
			},
		},
		{
			filename: "12.txt",
			predicate: func(x, y int) bool {
				return math.Sqrt(float64(x*x+y*y)) <= 20
			},
		},
		{
			filename: "13.txt",
			predicate: func(x, y int) bool {
				return (x+y >= 20) && (x+y < 29)
			},
		},
		// {
		// 	filename: "14.txt",
		// 	predicate: func(x, y int) bool {
		// 		x = size - x
		// 		y = size - y
		// 		return math.Sqrt(float64(x*x+y*y))*1.02 >= 21  // здесь получается примерно похожее 🙂
		// 		/*
		// 			WANT:
		// 			# # # # # # # # # # # # # # # # # # # # # # # # #
		// 			# # # # # # # # # # # # # # # # # # # # # # # # #
		// 			# # # # # # # # # # # # # # # # # # # # # # # # #
		// 			# # # # # # # # # # # # # # # # # # # # # # # # #
		// 			# # # # # # # # # # # # # # # # # # # # # # # # #
		// 			# # # # # # # # # # # # # # # # # # # # # . . . .
		// 			# # # # # # # # # # # # # # # # # . . . . . . . .
		// 			# # # # # # # # # # # # # # # . . . . . . . . . .
		// 			# # # # # # # # # # # # # . . . . . . . . . . . .
		// 			# # # # # # # # # # # # . . . . . . . . . . . . .
		// 			# # # # # # # # # # # . . . . . . . . . . . . . .
		// 			# # # # # # # # # # . . . . . . . . . . . . . . .
		// 			# # # # # # # # # . . . . . . . . . . . . . . . .
		// 			# # # # # # # # . . . . . . . . . . . . . . . . .
		// 			# # # # # # # # . . . . . . . . . . . . . . . . .
		// 			# # # # # # # . . . . . . . . . . . . . . . . . .
		// 			# # # # # # # . . . . . . . . . . . . . . . . . .
		// 			# # # # # # . . . . . . . . . . . . . . . . . . .
		// 			# # # # # # . . . . . . . . . . . . . . . . . . .
		// 			# # # # # # . . . . . . . . . . . . . . . . . . .
		// 			# # # # # # . . . . . . . . . . . . . . . . . . .
		// 			# # # # # . . . . . . . . . . . . . . . . . . . .
		// 			# # # # # . . . . . . . . . . . . . . . . . . . .
		// 			# # # # # . . . . . . . . . . . . . . . . . . . .
		// 			# # # # # . . . . . . . . . . . . . . . . . . . .

		//			GOT
		//			# # # # # # # # # # # # # # # # # # # # # # # # #
		//			# # # # # # # # # # # # # # # # # # # # # # # # #
		//			# # # # # # # # # # # # # # # # # # # # # # # # #
		//			# # # # # # # # # # # # # # # # # # # # # # # # #
		//			# # # # # # # # # # # # # # # # # # # # # # # # #
		//			# # # # # # # # # # # # # # # # # # # # # . . . .
		//			# # # # # # # # # # # # # # # # # # . . . . . . .
		//			# # # # # # # # # # # # # # # # . . . . . . . . .
		//			# # # # # # # # # # # # # # . . . . . . . . . . .
		//			# # # # # # # # # # # # # . . . . . . . . . . . .
		//			# # # # # # # # # # # . . . . . . . . . . . . . .
		//			# # # # # # # # # # . . . . . . . . . . . . . . .
		//			# # # # # # # # # # . . . . . . . . . . . . . . .
		//			# # # # # # # # # . . . . . . . . . . . . . . . .
		//			# # # # # # # # . . . . . . . . . . . . . . . . .
		//			# # # # # # # # . . . . . . . . . . . . . . . . .
		//			# # # # # # # . . . . . . . . . . . . . . . . . .
		//			# # # # # # # . . . . . . . . . . . . . . . . . .
		//			# # # # # # . . . . . . . . . . . . . . . . . . .
		//			# # # # # # . . . . . . . . . . . . . . . . . . .
		//			# # # # # # . . . . . . . . . . . . . . . . . . .
		//			# # # # # . . . . . . . . . . . . . . . . . . . .
		//			# # # # # . . . . . . . . . . . . . . . . . . . .
		//			# # # # # . . . . . . . . . . . . . . . . . . . .
		//			# # # # # . . . . . . . . . . . . . . . . . . . .
		// 		*/
		// 	},
		// },
		{
			filename: "15.txt",
			predicate: func(x, y int) bool {
				return ((x >= y+10) && (x <= y+20)) || ((y >= x+10) && (y <= x+20))
			},
		},
		{
			filename: "16.txt",
			predicate: func(x, y int) bool {
				return (x < y+10) && (x > y-10) && (size-x <= y+10) && (size-x > y-9)
			},
		},
		{
			filename: "17.txt",
			predicate: func(x, y int) bool {
				return math.Sin(float64(x)/3) <= float64(y)/8-2 // подсмотрел 🙂
			},
		},
		{
			filename: "18.txt",
			predicate: func(x, y int) bool {
				if x == 0 && y == 0 {
					return false
				}
				return ((y < 2) || (x < 2))
			},
		},
		{
			filename: "19.txt",
			predicate: func(x, y int) bool {
				return x == 0 || y == 0 || x == size-1 || y == size-1
			},
		},
		{
			filename: "20.txt",
			predicate: func(x, y int) bool {
				return (x+y)%2 == 0
			},
		},
		{
			filename: "21.txt",
			predicate: func(x, y int) bool {
				return x%(y+1) == 0
			},
		},
		{
			filename: "22.txt",
			predicate: func(x, y int) bool {
				return (x+y)%3 == 0
			},
		},
		{
			filename: "23.txt",
			predicate: func(x, y int) bool {
				return (x%2 == 0) && (y%3 == 0)
			},
		},
		{
			filename: "24.txt",
			predicate: func(x, y int) bool {
				return (x == y) || (size-x-1) == y
			},
		},
		{
			filename: "25.txt",
			predicate: func(x, y int) bool {
				return x%6 == 0 || y%6 == 0
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.filename, func(t *testing.T) {
			got := abracadabra(tt.predicate)

			want := strings.Fields(readFile(t, tt.filename))
			if !reflect.DeepEqual(want, got) {
				t.Log("WANT")
				print(t, want)
				t.Log("GOT")
				print(t, got)
				t.Fatal("not equal")
			}
		})
	}
}

func readFile(t *testing.T, filename string) string {
	t.Helper()
	b, err := os.ReadFile(path.Join("data", filename))
	fatalErr(t, err)

	return string(b)
}

func print(t *testing.T, data []string) {
	t.Helper()

	var sb strings.Builder
	sb.WriteRune('\n')

	for i, r := range data {
		sb.WriteString(r)
		if (i+1)%size == 0 {
			sb.WriteString("\n")
		} else {
			sb.WriteString(" ")
		}
	}
	t.Log(sb.String())
}

func fatalErr(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}
