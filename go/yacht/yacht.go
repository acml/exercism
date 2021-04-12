package yacht

// Score calculates the score of a throw of the dice which depends on category
// chosen.
func Score(dice []int, category string) (score int) {
	switch category {
	case "ones":
		return count(dice, 1)
	case "twos":
		return count(dice, 2) * 2
	case "threes":
		return count(dice, 3) * 3
	case "fours":
		return count(dice, 4) * 4
	case "fives":
		return count(dice, 5) * 5
	case "sixes":
		return count(dice, 6) * 6
	case "yacht":
		return yacht(dice)
	case "choice":
		return choice(dice)
	case "full house":
		return fullHouse(dice)
	case "four of a kind":
		return fourOfAKind(dice)
	case "little straight":
		return straight(dice, 1)
	default: //case "big straight":
		return straight(dice, 2)
	}
}

func count(dice []int, v int) (n int) {
	for _, d := range dice {
		if d == v {
			n++
		}
	}
	return n
}

func diceCounts(dice []int) map[int]int {
	counts := map[int]int{}
	for _, d := range dice {
		counts[d]++
	}
	return counts
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
	counts := diceCounts(dice)
	if len(counts) != 2 {
		return 0
	}

	for d, count := range counts {
		if count != 2 && count != 3 {
			return 0
		}
		score += d * count
	}
	return score
}

func fourOfAKind(dice []int) int {
	counts := diceCounts(dice)
	if len(counts) > 2 {
		return 0
	}

	for d, count := range counts {
		if count >= 4 {
			return d * 4
		}
	}
	return 0
}

func straight(dice []int, offset int) int {
	counts := diceCounts(dice)
	if len(counts) < 5 {
		return 0
	}

	for d := offset; d <= offset+4; d++ {
		if counts[d] != 1 {
			return 0
		}
	}
	return 30
}
