package bitboard

func Rook(pos int) (int, uint64) {
	bMoves := lines(pos)

	return countMoves(bMoves), bMoves
}
