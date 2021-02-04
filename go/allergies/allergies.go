// Package allergies provides functions which given a person's allergy score,
// determine whether or not they're allergic to a given item, and their full
// list of allergies.
package allergies

const (
	eggs uint = 1 << iota
	peanuts
	shellfish
	strawberries
	tomatoes
	chocolate
	pollen
	cats
)

var allergies = map[uint]string{
	eggs:         "eggs",
	peanuts:      "peanuts",
	shellfish:    "shellfish",
	strawberries: "strawberries",
	tomatoes:     "tomatoes",
	chocolate:    "chocolate",
	pollen:       "pollen",
	cats:         "cats",
}

// Allergies given a person's allergy score, returns their full list of
// allergies.
func Allergies(score uint) []string {
	res := []string{}
	for i := 0; i < len(allergies); i++ {
		allergy, ok := allergies[(1<<i)&score]
		if ok {
			res = append(res, allergy)
		}
	}
	return res
}

// AllergicTo given a person's allergy score, determines whether or not they're
// allergic to a given item.
func AllergicTo(score uint, allergen string) bool {
	for k, v := range allergies {
		if (score&k) > 0 && v == allergen {
			return true
		}
	}
	return false
}
