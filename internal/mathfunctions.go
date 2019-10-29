package internal

func (i *Integer) IsPrime() bool {
	for a := 2; a < i.n; a++ {
		if i.n%a == 0 {
			return false
		}
	}
	return true
}

func (i *Integer64) Factors() []int64 {
	var retval = []int64{}

	var mid int64 = 0
	if i.n%2 == 0 {
		mid = i.n / 2
	} else {
		mid = (i.n + 1) / 2
	}

	var x int64
	for x = 2; x < mid; x++ {
		if mid%x == 0 {
			var f int64 = mid / x

			i.n = x
			if !i.Contains(retval) {
				retval = append(retval, x)
			}

			i.n = f
			if f != x && !i.Contains(retval) {
				retval = append(retval, f)
			}
		}
	}

	return retval
}
