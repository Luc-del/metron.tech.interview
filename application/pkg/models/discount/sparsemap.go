package discount

// sparsePositiveIntMap is a sparse map which keys are positive integer
// it allows to store not consecutive keys and to retrieve the exact or lowest value
type sparsePositiveIntMap struct {
	values map[int]int
}

// getEqualOrLower returns the value associated to the key
// If absent, returns the value of the closest lower key
// If none, returns 0
func (s sparsePositiveIntMap) getEqualOrLower(key int) int {
	for ; key >= 0; key-- {
		if val, ok := s.values[key]; ok {
			return val
		}
	}

	return 0
}

func (s sparsePositiveIntMap) setValue(k, v int) {
	s.values[k] = v
}

func (s sparsePositiveIntMap) len() int {
	return len(s.values)
}

func newEmpty(capacity int) sparsePositiveIntMap {
	return sparsePositiveIntMap{
		values: make(map[int]int, capacity),
	}
}

func (s sparsePositiveIntMap) copy() sparsePositiveIntMap {
	newMap := newEmpty(s.len())
	for k, v := range s.values {
		newMap.values[k] = v
	}

	return newMap
}
