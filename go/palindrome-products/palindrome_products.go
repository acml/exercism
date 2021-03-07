package palindrome

import "errors"

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

	products := map[int]struct{}{}
	for i := fmin; i <= fmax; i++ {
		for j := fmin; j <= i; j++ {
			product := i * j
			if !isPalindrome(product) {
				continue
			}

			if _, ok := products[product]; !ok {
				products[product] = struct{}{}
			}
		}
	}

	for product := range products {
		if pmin.Product == 0 || product < pmin.Product {
			pmin = Product{
				Product:        product,
				Factorizations: factors(product, fmin, fmax),
			}
		}

		if pmax.Product == 0 || product > pmax.Product {
			pmax = Product{
				Product:        product,
				Factorizations: factors(product, fmin, fmax),
			}
		}
	}

	if pmin.Product == 0 && pmax.Product == 0 {
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

func factors(n, min, max int) (factors [][2]int) {
	for i := min; i <= max; i++ {
		for j := max; j >= i; j-- {
			if n > 1 && i*j == n {
				factors = append(factors, [2]int{i, j})
			}
		}
	}
	return factors
}
