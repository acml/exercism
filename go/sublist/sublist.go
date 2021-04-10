package sublist

// Relation of two lists
type Relation string

// Sublist determines the relation of two given lists.
func Sublist(listOne, listTwo []int) Relation {

	if len(listOne) == len(listTwo) {
		for p, v := range listOne {
			if v != listTwo[p] {
				return "unequal"
			}
		}
		return "equal"
	}

	var possibleResult Relation = "sublist"
	var listA, listB []int
	if len(listOne) < len(listTwo) {
		listA = listOne
		listB = listTwo

	} else {
		listA = listTwo
		listB = listOne
		possibleResult = "superlist"
	}

	if len(listA) == 0 {
		return possibleResult
	}

	for pb, vb := range listB {
		if len(listB)-pb < len(listA) {
			return "unequal"
		}
		if listA[0] == vb && listMatch(listA, listB[pb:]) {
			return possibleResult
		}
	}
	return "unequal"
}

func listMatch(listOne, listTwo []int) bool {
	for p, v := range listOne {
		if v != listTwo[p] {
			return false
		}
	}
	return true
}
