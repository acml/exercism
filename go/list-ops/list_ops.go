package listops

type IntList []int
type binFunc func(int, int) int
type predFunc func(int) bool
type unaryFunc func(x int) int

func (s IntList) Length() int {
	n := 0
	for _, v := range s {
		_ = v
		n++
	}
	return n
}

func (s1 IntList) Append(s2 IntList) IntList {
	s := make(IntList, s1.Length() + s2.Length())
	for i, v := range s1 {
		s[i] = v
	}
	for i, v := range s2 {
		s[i+s1.Length()] = v
	}
	return s
}

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

func (s IntList) Map(f func(int) int) IntList {
	r := make(IntList, s.Length())
	for i, v := range s {
		r[i] = f(v)
	}
	return r
}

func (s IntList) Foldl(f func(int, int) int, acc int) int {
	for _, v := range s {
		acc = f(acc, v)
	}
	return acc
}

func (s IntList) Foldr(f func(int, int) int, acc int) int {
	for i := s.Length() - 1; i >= 0; i-- {
		acc = f(s[i], acc)
	}
	return acc
}

func (s IntList) Reverse() IntList {
	r := make(IntList, s.Length())
	for i, j := s.Length() - 1, 0; i >= 0; {
		r[i] = s[j]
		i--
		j++
	}
	return r
}
