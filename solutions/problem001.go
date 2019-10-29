package solutions

/**
 * If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
 *
 * Find the sum of all the multiples of 3 or 5 below 1000.
 */

// Problem is the base receiver struct
type Problem struct{}

// Problem001 - Multiples of 3 and 5
func (p *Problem) Problem001() int {
	const limit = 1000
	retval := 0

	// iterate over range looking for multiples of 3 or 5
	// this problem doesn't tak into consideration those values that
	// have both, 3 and 5 as multiples like 15
	for i := 0; i < limit; i++ {
		if i%3 == 0 {
			retval += i
		} else if i%5 == 0 {
			retval += i
		}
	}

	return retval
}
