package bitboard

func lines(pos int) uint64 {
	xL, xR, yD, yU := coordDeltas(pos)

	var bPos uint64 = 1 << pos
	var bMoves uint64
	bMoves += _lines(bPos, lineLeft, xL)
	bMoves += _lines(bPos, lineRight, xR)
	bMoves += _lines(bPos, lineUp, yU)
	bMoves += _lines(bPos, lineDown, yD)

	return bMoves
}

type lineBitFunc func(bPos uint64, d int) uint64

func lineLeft(bPos uint64, d int) uint64 {
	return bPos >> d
}

func lineRight(bPos uint64, d int) uint64 {
	return bPos << d
}

func lineUp(bPos uint64, d int) uint64 {
	return bPos << (d * 8)
}

func lineDown(bPos uint64, d int) uint64 {
	return bPos >> (d * 8)
}

func _lines(bPos uint64, bitFunc lineBitFunc, d int) uint64 {
	var bMoves uint64
	for d > 0 {
		bMoves |= bitFunc(bPos, d)
		d--
	}
	return bMoves
}

func cross(pos int) uint64 {
	xL, xR, yD, yU := coordDeltas(pos)
	var bPos uint64 = 1 << pos
	var bMoves uint64

	bMoves += _cross(bPos, crossLeftTop, xL, yU)
	bMoves += _cross(bPos, crossRightTop, xR, yU)
	bMoves += _cross(bPos, crossRightBottom, xR, yD)
	bMoves += _cross(bPos, crossLeftBottom, xL, yD)
	return bMoves
}

type crossBitFunc func(bPos uint64, x, y int) uint64

func crossLeftTop(bPos uint64, x, y int) uint64 {
	return bPos << (8*y - x)
}

func crossRightTop(bPos uint64, x, y int) uint64 {
	return bPos << (8*y + x)
}

func crossRightBottom(bPos uint64, x, y int) uint64 {
	return bPos >> (8*y - x)
}

func crossLeftBottom(bPos uint64, x, y int) uint64 {
	return bPos >> (8*y + x)
}

func _cross(bPos uint64, bFunc crossBitFunc, dX, dY int) uint64 {
	x := 1
	y := 1
	var bMoves uint64
	for x <= dX && y <= dY {
		bMoves |= bFunc(bPos, x, y)
		x++
		y++
	}

	return bMoves
}
