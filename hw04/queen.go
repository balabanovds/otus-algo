package bitboard

func Queen(pos int) (int, uint64) {
	bMoves := cross(pos)
	bMoves += lines(pos)

	return countMoves(bMoves), bMoves
}
