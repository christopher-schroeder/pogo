package pogo

import (
	"fmt"
	"math/rand"
)

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


func Equaty(hand []uint8, board []uint8) {
	//var mask uint64 = (1 << hand[0]) | (1 << hand[1])

	var allcards [52]uint8 = [52]uint8 {
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12,
		13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25,
		26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38,
		39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	}

	for i:=0; i<len(board); i++ {
		allcards[board[i]] = 255
	}

	last_index := 51
	for i:=51; i>=0; i-- {
		if allcards[i] == 255 {
			allcards[last_index], allcards[i] = allcards[i], allcards[last_index]
			last_index--
		}
	}

	for i:=0; i<1000; i++ {
		for j:=last_index; j>last_index-5+len(board); j-- {
			fmt.Print(allcards[rand.Intn(j)], ",")
		}
		fmt.Println("")
		fmt.Println("")
	}

	//board_copy := make([]uint8, len(board))

}