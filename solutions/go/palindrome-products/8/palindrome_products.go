package palindrome

import "errors"

const (
	maxInt = int(^uint(0) >> 1)
	minInt = -maxInt - 1
)

// Product stores palindromes with factors.
type Product struct {
	Product        int      // palindromic, of course
	Factorizations [][2]int //list of all possible two-factor factorizations of Product, within given limits, in order
}

// Products returns the largest and smallest palindromes, along with the factors
// of each within the range.
func Products(fmin, fmax int) (pmin, pmax Product, err error) {

	if fmin > fmax {
		return pmin, pmax, errors.New("fmin > fmax")
	}

	pmin.Product = maxInt
	pmax.Product = minInt
	for i := fmin; i <= fmax; i++ {
		for j := i; j <= fmax; j++ {
			p := i * j
			if pmin.Product < p && p < pmax.Product {
				break
			}
			switch {
			case p == pmin.Product:
				pmin.Factorizations = append(pmin.Factorizations, [2]int{i, j})
			case p == pmax.Product:
				pmax.Factorizations = append(pmax.Factorizations, [2]int{i, j})
			case (p < pmin.Product || p > pmax.Product) && isPalindrome(p):
				switch {
				case p < pmin.Product:
					pmin.Product = p
					pmin.Factorizations = append(pmin.Factorizations[:0], [2]int{i, j})
				case p > pmax.Product:
					pmax.Product = p
					pmax.Factorizations = append(pmax.Factorizations[:0], [2]int{i, j})
				}
			}
		}
	}

	if pmin.Product == maxInt {
		return pmin, pmax, errors.New("no palindromes")
	}

	return pmin, pmax, nil
}

func isPalindrome(number int) bool {
	r := 0
	for n := number; n > 0; n /= 10 {
		r = 10*r + n%10
	}
	return number == r
}
