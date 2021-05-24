package bitboard

func Bishop(pos int) (int, uint64) {
	bMoves := cross(pos)

	return countMoves(bMoves), bMoves
}
