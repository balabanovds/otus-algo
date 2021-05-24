package bitboard

func King(pos int) (int, uint64) {
	var maskA uint64 = 0b11111110_11111110_11111110_11111110_11111110_11111110_11111110_11111110
	var maskH uint64 = 0b01111111_01111111_01111111_01111111_01111111_01111111_01111111_01111111

	var bPos uint64 = 1 << pos                      // b2
	bMoves := maskH & (bPos>>9 | bPos>>1 | bPos<<7) // a1,a2,a3
	bMoves |= maskA & (bPos>>7 | bPos<<1 | bPos<<9) // c1,c2,c3
	bMoves |= bPos>>8 | bPos<<8                     // b1,b4

	return countMoves(bMoves), bMoves
}
