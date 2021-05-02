package hw01

const size = 25

func abracadabra(pred func(x, y int) bool) []string {
	res := make([]string, 0, size*size)

	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			if pred(x, y) {
				res = append(res, "#")
			} else {
				res = append(res, ".")
			}
		}
	}

	return res
}
