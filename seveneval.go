package pogo

func GetRank(i uint8, j uint8, k uint8, m uint8, n uint8, p uint8, q uint8) uint16 {
	key := card[i] + card[j] + card[k] + card[m] + card[n] + card[p] + card[q];

	suit := flush_check[key >> FLUSH_BIT_SHIFT];
	if (NOT_A_SUIT != suit) {
		s := suit_kronecker[suit];
		return flush_ranks[s[i] | s[j] | s[k] | s[m] | s[n] | s[p] | s[q]]
	}

	hash := FACE_BIT_MASK & uint32(31 * uint64(key));

	return rank_hash[offsets[hash >> RANK_OFFSET_SHIFT] + uint16(hash & RANK_HASH_MOD)]
}