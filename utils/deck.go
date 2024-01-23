package utils

var fullDeck = []int{
	0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
	21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39,
}

var miniDeck = []int{10, 20, 11, 21, 14, 24, 15, 25, 16, 26, 17, 27, 18, 28}

var MicroDeckNoFlor = []int{20, 0, 26, 36, 12, 16, 5}

func Deck(size int) []int {
	if size < 7 {
		panic("deck size cannot be lower than 1+3+3=7")
	}

	if size <= 14 {
		return miniDeck[:size]
	}

	return fullDeck[:size]
}
