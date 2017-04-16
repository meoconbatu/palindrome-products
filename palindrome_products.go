package palindrome

import (
	"errors"
)

const testVersion = 1

type Product struct {
	Product        int      // palindromic, of course
	Factorizations [][2]int //list of all possible two-factor factorizations of Product, within given limits, in order
}

func Products(fmin, fmax int) (pmin, pmax Product, err error) {
	if err = ValidInput(fmin, fmax); err != nil {
		return pmin, pmax, err
	}
	pmin.Product, pmax.Product = fmax*fmax, 0
	for i := fmin; i <= fmax; i++ {
		for j := i; j <= fmax; j++ {
			product := i * j
			if product < 10 || product == Reverse(product) {
				if product > pmax.Product {
					pmax.Product = product
					pmax.Factorizations = [][2]int{{i, j}}
				} else if product == pmax.Product {
					pmax.Factorizations = append(pmax.Factorizations, [2]int{i, j})
				}
				if product < pmin.Product {
					pmin.Product = product
					pmin.Factorizations = [][2]int{{i, j}}
				} else if product == pmin.Product {
					pmin.Factorizations = append(pmin.Factorizations, [2]int{i, j})
				}
			}
		}
	}
	if pmin.Factorizations == nil && pmax.Factorizations == nil {
		err = errors.New("No palindromes")
	}
	return pmin, pmax, err
}

func ValidInput(fmin, fmax int) error {
	if fmax < fmin {
		return errors.New("fmin > fmax")
	}
	return nil
}
func Reverse(n int) int {
	new_int := 0
	for n > 0 {
		remainder := n % 10
		new_int *= 10
		new_int += remainder
		n /= 10
	}
	return new_int
}
