package harry

import (
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
				return (size - x - 1) > y-6
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
