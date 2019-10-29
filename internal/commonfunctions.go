package internal

type Integer struct {
	n int
}
type Integer64 struct {
	n int64
}

// Contains tests each element in the array for t
func (i *Integer) Contains(array []int) bool {
	if len(array) == 0 {
		return false
	}

	for _, x := range array {
		if x == i.n {
			return true
		}
	}

	return false
}

func (i *Integer64) Contains(array []int64) bool {
	if len(array) == 0 {
		return false
	}

	for _, x := range array {
		if x == i.n {
			return true
		}
	}

	return false
}
