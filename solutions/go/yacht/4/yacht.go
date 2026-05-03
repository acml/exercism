package yacht

var scoreFunc = map[string]func([]int) int{
	"ones":            ones,
	"twos":            twos,
	"threes":          threes,
	"fours":           fours,
	"fives":           fives,
	"sixes":           sixes,
	"yacht":           yacht,
	"choice":          choice,
	"full house":      fullHouse,
	"four of a kind":  fourOfAKind,
	"little straight": littleStraight,
	"big straight":    bigStraight,
}

// Score calculates the score of a throw of the dice which depends on category
// chosen.
func Score(dice []int, category string) (score int) {
	return scoreFunc[category](dice)
}

func count(dice []int, v int) (n int) {
	for _, d := range dice {
		if d == v {
			n++
		}
	}
	return n
}

func diceMap(dice []int) map[int]int {
	diceCounts := map[int]int{}
	for _, d := range dice {
		diceCounts[d]++
	}
	return diceCounts
}

func ones(dice []int) int {
	return count(dice, 1)
}

func twos(dice []int) int {
	return 2 * count(dice, 2)
}

func threes(dice []int) int {
	return 3 * count(dice, 3)
}

func fours(dice []int) int {
	return 4 * count(dice, 4)
}

func fives(dice []int) int {
	return 5 * count(dice, 5)
}

func sixes(dice []int) int {
	return 6 * count(dice, 6)
}

func yacht(dice []int) int {
	for p, d := range dice {
		if p > 0 && d != dice[p-1] {
			return 0
		}
	}
	return 50
}

func choice(dice []int) (score int) {
	for _, d := range dice {
		score += d
	}
	return score
}

func fullHouse(dice []int) (score int) {
	diceCounts := diceMap(dice)
	if len(diceCounts) != 2 {
		return 0
	}

	for d, count := range diceCounts {
		if count != 2 && count != 3 {
			return 0
		}
		score += d * count
	}
	return score
}

func fourOfAKind(dice []int) int {
	diceCounts := diceMap(dice)
	if len(diceCounts) > 2 {
		return 0
	}

	for d, count := range diceCounts {
		if count >= 4 {
			return d * 4
		}
	}
	return 0
}

func littleStraight(dice []int) int {
	return straight(dice, 1)
}

func bigStraight(dice []int) int {
	return straight(dice, 2)
}

func straight(dice []int, offset int) int {
	diceCounts := diceMap(dice)
	if len(diceCounts) < 5 {
		return 0
	}

	for d := offset; d <= offset+4; d++ {
		if diceCounts[d] != 1 {
			return 0
		}
	}
	return 30
}
