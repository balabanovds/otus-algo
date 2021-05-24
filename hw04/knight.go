package bitboard

func Knight(pos int) (int, uint64) {
	var maskA uint64 = 0b11111110_11111110_11111110_11111110_11111110_11111110_11111110_11111110
	var maskAB uint64 = 0b11111100_11111100_11111100_11111100_11111100_11111100_11111100_11111100
	var maskH uint64 = 0b01111111_01111111_01111111_01111111_01111111_01111111_01111111_01111111
	var maskGH uint64 = 0b00111111_00111111_00111111_00111111_00111111_00111111_00111111_00111111

	var bPos uint64 = 1 << pos              // c3
	bMoves := maskGH & (bPos>>10 | bPos<<6) // a2, a4
	bMoves |= maskH & (bPos>>17 | bPos<<15) // b1, b5
	bMoves |= maskA & (bPos>>15 | bPos<<17) // d3, d5
	bMoves |= maskAB & (bPos>>6 | bPos<<10) // e2, e4

	return countMoves(bMoves), bMoves
}
