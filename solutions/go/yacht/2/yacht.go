package yacht

var weight = map[string]int{
	"ones":   1,
	"twos":   2,
	"threes": 3,
	"fours":  4,
	"fives":  5,
	"sixes":  6,
}

// Score calculates the score of a throw of the dice which depends on category
// chosen.
func Score(dice []int, category string) (score int) {
	switch category {
	case "ones", "twos", "threes", "fours", "fives", "sixes":
		v := weight[category]
		for _, d := range dice {
			if d == v {
				score += v
			}
		}
		return score
	case "yacht":
		for p, d := range dice {
			if p > 0 && d != dice[p-1] {
				return 0
			}
		}
		return 50
	case "choice":
		for _, d := range dice {
			score += d
		}
		return score
	}

	diceCounts := map[int]int{}
	for _, d := range dice {
		diceCounts[d]++
	}

	switch category {
	case "full house":
		if len(diceCounts) != 2 {
			return 0
		}

		for d, count := range diceCounts {
			if count != 2 && count != 3 {
				return 0
			}
			score += d * count
		}
	case "four of a kind":
		if len(diceCounts) > 2 {
			return 0
		}

		for d, count := range diceCounts {
			if count < 4 {
				continue
			}
			score += d * 4
			break
		}
	case "little straight", "big straight":
		if len(diceCounts) < 5 {
			return 0
		}
		offset := 1
		if category == "big straight" {
			offset = 2
		}
		for i := 0; i < 5; i++ {
			if _, ok := diceCounts[i+offset]; !ok {
				return 0
			}
		}
		score = 30
	}
	return score
}
