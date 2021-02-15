package listops

// IntList represent a list of integers
type IntList []int
type binFunc func(int, int) int
type predFunc func(int) bool
type unaryFunc func(x int) int

// Length given a list, returns the total number of items within it.
func (s IntList) Length() int {
	n := 0
	for _, v := range s {
		_ = v
		n++
	}
	return n
}

// Append given two lists, adds all items in the second list to the end of the
// first list.
func (s IntList) Append(s2 IntList) IntList {
	res := make(IntList, s.Length()+s2.Length())
	for i, v := range s {
		res[i] = v
	}
	s1len := s.Length()
	for i, v := range s2 {
		res[s1len+i] = v
	}
	return res
}

// Concat given a series of lists, combines all items in all lists into one
// flattened list.
func (s IntList) Concat(sl []IntList) IntList {
	l := s.Length()
	for _, sli := range sl {
		l += sli.Length()
	}

	r := make(IntList, l)
	i := 0
	for _, si := range s {
		r[i] = si
		i++
	}
	for _, sli := range sl {
		for _, si := range sli {
			r[i] = si
			i++
		}
	}
	return r
}

// Filter given a predicate and a list, returns the list of all items for which
// `predicate(item)` is True.
func (s IntList) Filter(predicate func(int) bool) IntList {
	r := make(IntList, s.Length())
	i := 0
	for _, si := range s {
		if predicate(si) {
			r[i] = si
			i++
		}
	}
	return r[:i]
}

// Map given a function and a list, returns the list of the results of applying
// `function(item)` on all items.
func (s IntList) Map(f func(int) int) IntList {
	r := make(IntList, s.Length())
	for i, v := range s {
		r[i] = f(v)
	}
	return r
}

// Foldl given a function, a list, and initial accumulator, folds (reduces) each
// item into the accumulator from the left using `function(accumulator, item)`.
func (s IntList) Foldl(f func(int, int) int, acc int) int {
	for _, v := range s {
		acc = f(acc, v)
	}
	return acc
}

// Foldr given a function, a list, and an initial accumulator, folds (reduces)
// each item into the accumulator from the right using `function(item,
// accumulator)`.
func (s IntList) Foldr(f func(int, int) int, acc int) int {
	for i := s.Length() - 1; i >= 0; i-- {
		acc = f(s[i], acc)
	}
	return acc
}

// Reverse given a list, returns a list with all the original items, but in
// reversed order.
func (s IntList) Reverse() IntList {
	r := make(IntList, s.Length())
	for i, j := s.Length()-1, 0; i >= 0; {
		r[i] = s[j]
		i--
		j++
	}
	return r
}
