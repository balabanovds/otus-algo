package bitboard

func countMoves(bMoves uint64) int {
	cntr := 0
	for bMoves != 0 {
		bMoves &= bMoves - 1
		cntr++
	}
	return cntr
}
func coords(pos int) (x int, y int) {
	var bPos uint64 = 1 << pos

	for bPos > 0 {
		if bPos <= 128 {
			x = xPos(bPos)
		}
		y++
		bPos >>= 8
	}
	return
}

func coordDeltas(pos int) (xLeft int, xRight int, yDown int, yUp int) {
	x, y := coords(pos)
	xLeft = x - 1
	xRight = 8 - x
	yDown = y - 1
	yUp = 8 - y
	return
}

func xPos(num uint64) int {
	var pos, cntr uint64 = 1, 1
	for pos&num == 0 {
		pos <<= 1
		cntr++
	}
	return int(cntr)
}
